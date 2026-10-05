#!/usr/bin/env python3
"""
Merge redacted JSON skeletons (from test/integration HTTP logger NDJSON) into
openapi/openapi.yaml response schemas. Does not read or write secrets.

Uses **surgical line edits** so the rest of the YAML file is unchanged (no full-file
reformat). Input lines must be response records with json_skeleton (value-redacted).
Only status 200/201 are applied. Inserts sibling `x-structure-from-redacted-live-run`
under `application/json` when missing.

Usage:
  python utils/merge_skeleton_into_openapi.py [run.jsonl] [openapi/openapi.yaml]

Broader capture (all GET operations, anonymized) before merging:
  python utils/collect_live_response_shapes.py --truncate --out test/integration/_live_shapes/sweep.jsonl
  # Append integration test logs, then:
  python utils/merge_skeleton_into_openapi.py test/integration/_live_shapes/sweep.jsonl openapi/openapi.yaml
"""

from __future__ import annotations

import io
import json
import re
import sys
from collections.abc import Mapping, Sequence
from dataclasses import dataclass
from pathlib import Path

try:
    from ruamel.yaml import YAML
except ImportError:
    print("Install ruamel.yaml: pip install ruamel.yaml", file=sys.stderr)
    sys.exit(1)


def merge_skeleton_union(a: object, b: object) -> object:
    """Union-merge two skeleton trees (same contract as Go jsonSkeleton output)."""
    if a is None:
        return b
    if b is None:
        return a
    if isinstance(a, Mapping) and isinstance(b, Mapping):
        keys = set(a) | set(b)
        return {
            k: merge_skeleton_union(a.get(k), b.get(k))
            if k in a and k in b
            else (a[k] if k in a else b[k])
            for k in keys
        }
    if isinstance(a, Sequence) and not isinstance(a, (str, bytes)) and isinstance(
        b, Sequence
    ) and not isinstance(b, (str, bytes)):
        if len(a) == 0:
            return list(b)
        if len(b) == 0:
            return list(a)
        return [merge_skeleton_union(a[0], b[0])]
    return a


def skeleton_to_schema(obj: object) -> dict:
    """Turn a redacted skeleton into an OpenAPI 3.0 schema object."""
    if obj is None:
        return {"type": "string", "nullable": True}
    if isinstance(obj, bool):
        return {"type": "boolean"}
    if isinstance(obj, (int, float)):
        return {"type": "number"}
    if isinstance(obj, str):
        return {"type": "string"}
    if isinstance(obj, list):
        if len(obj) == 0:
            return {"type": "array", "items": {"type": "object", "additionalProperties": True}}
        return {"type": "array", "items": skeleton_to_schema(obj[0])}
    if isinstance(obj, dict):
        if obj.get("_non_json") or obj.get("_parse_error"):
            return {"type": "object", "additionalProperties": True}
        if "_type" in obj:
            return {"type": "object", "additionalProperties": True}
        props = {k: skeleton_to_schema(v) for k, v in sorted(obj.items())}
        return {"type": "object", "properties": props}
    return {"type": "object", "additionalProperties": True}


def load_ndjson(path: Path) -> list[dict]:
    rows: list[dict] = []
    for line in path.read_text(encoding="utf-8-sig").splitlines():
        line = line.strip()
        if not line:
            continue
        rows.append(json.loads(line))
    return rows


def aggregate_skeletons(rows: list[dict]) -> dict[tuple[str, str], object]:
    """(method, path) -> merged json_skeleton."""
    out: dict[tuple[str, str], object | None] = {}
    for row in rows:
        if row.get("kind") != "response":
            continue
        status = int(row.get("status", 0))
        if status not in (200, 201):
            continue
        sk = row.get("json_skeleton")
        if sk is None:
            continue
        method = str(row.get("method", "GET")).upper()
        path = str(row.get("path", ""))
        key = (method, path)
        if key not in out or out[key] is None:
            out[key] = sk
        else:
            out[key] = merge_skeleton_union(out[key], sk)
    return {k: v for k, v in out.items() if v is not None}


def indent_level(line: str) -> int:
    return len(line) - len(line.lstrip(" "))


def find_path_line(lines: list[str], path: str) -> int | None:
    needle = f"  {path}:"
    for i, line in enumerate(lines):
        if line.rstrip() == needle:
            return i
    return None


def find_method_line(lines: list[str], path_idx: int, method: str) -> int | None:
    m = f"    {method.lower()}:"
    for i in range(path_idx + 1, len(lines)):
        line = lines[i]
        if line.startswith("  /v1/") and indent_level(line) == 2:
            break
        if line.rstrip() == m:
            return i
    return None


_METHOD_SIBLING_KEYS = frozenset(
    (
        "get:",
        "post:",
        "put:",
        "patch:",
        "delete:",
        "options:",
        "head:",
        "trace:",
        "parameters:",
    )
)


def method_operation_block_end(lines: list[str], method_idx: int) -> int:
    """First line index after this operation: next path or sibling verb/parameters at same indent."""
    base = indent_level(lines[method_idx])
    n = len(lines)
    i = method_idx + 1
    while i < n:
        line = lines[i]
        if line.startswith("  /v1/") and indent_level(line) == 2:
            return i
        if indent_level(line) == base and line.strip():
            key = line.strip().split(None, 1)[0].lower()
            if key in _METHOD_SIBLING_KEYS:
                return i
        i += 1
    return n


def find_schema_block_range(
    lines: list[str], method_idx: int, status_code: str
) -> tuple[int, int] | None:
    """
    Return (schema_key_line_index, end_exclusive) where [schema_key+1, end_exclusive)
    is the body under the schema mapping (children only).
    """
    code_tag = f"'{status_code}':"
    code_tag2 = f'"{status_code}":'
    end_m = method_operation_block_end(lines, method_idx)
    i = method_idx + 1
    while i < end_m:
        line = lines[i]
        if line.startswith("  /v1/") and indent_level(line) == 2:
            return None
        if code_tag in line or code_tag2 in line:
            # Search forward for application/json -> schema within this status block only.
            # If there is no schema (e.g. application/json: {}), return None — do not keep
            # scanning forward or we will match the next HTTP method's response under the same path.
            j = i + 1
            seen_json = False
            while j < end_m:
                lj = lines[j]
                if indent_level(lj) <= indent_level(line) and lj.strip() and j > i:
                    # Left the '200'/'201' block without finding schema
                    return None
                if "application/json:" in lj:
                    seen_json = True
                if seen_json and lj.strip().startswith("schema:"):
                    base = indent_level(lj)
                    end = j + 1
                    while end < end_m:
                        if not lines[end].strip():
                            end += 1
                            continue
                        if indent_level(lines[end]) <= base:
                            return j, end
                        end += 1
                    return j, end
                j += 1
            return None
        i += 1
    return None


_INLINE_JSON_EMPTY = re.compile(r"^\s*application/json:\s*\{\}\s*$")


def find_inline_application_json_empty(
    lines: list[str], method_idx: int, status_code: str
) -> int | None:
    """Line index of `application/json: {}` under responses -> status (no schema key)."""
    code_tag = f"'{status_code}':"
    code_tag2 = f'"{status_code}":'
    n = method_operation_block_end(lines, method_idx)
    i = method_idx + 1
    while i < n:
        line = lines[i]
        if line.startswith("  /v1/") and indent_level(line) == 2:
            return None
        if code_tag in line or code_tag2 in line:
            j = i + 1
            while j < n:
                lj = lines[j]
                if indent_level(lj) <= indent_level(line) and lj.strip() and j > i:
                    return None
                if _INLINE_JSON_EMPTY.match(lj):
                    return j
                j += 1
            return None
        i += 1
    return None


@dataclass(frozen=True)
class SchemaPatch:
    """Replace children under an existing `schema:` key."""

    schema_key_idx: int
    end_excl: int
    new_body: list[str]
    marker_base_indent: int


@dataclass(frozen=True)
class InlineJsonPatch:
    """Replace `application/json: {}` with schema + marker."""

    app_json_line_idx: int
    new_body: list[str]


def schema_dict_to_body_lines(schema: dict, body_indent: int) -> list[str]:
    """Serialize schema to YAML lines with fixed indent (spaces before each line)."""
    y = YAML()
    y.default_flow_style = False
    buf = io.StringIO()
    y.dump(schema, buf)
    raw = buf.getvalue().rstrip("\n")
    if not raw:
        return [" " * body_indent + "{}"]
    out: list[str] = []
    for ln in raw.splitlines():
        out.append(" " * body_indent + ln)
    return out


def surgical_patch_openapi_text(lines: list[str], aggregated: dict[tuple[str, str], object]) -> tuple[list[str], int]:
    """Return (new_lines, updated_count).

    Patches are collected using stable line indices on the **original** text, then applied
    from the **bottom of the file upward** so earlier indices stay valid.
    """
    marker_key = "x-structure-from-redacted-live-run"
    schema_patches: list[SchemaPatch] = []
    inline_patches: list[InlineJsonPatch] = []

    for (method, path), skel in aggregated.items():
        pi = find_path_line(lines, path)
        if pi is None:
            continue
        mi = find_method_line(lines, pi, method)
        if mi is None:
            continue
        new_schema = skeleton_to_schema(skel)
        placed = False
        for code in ("200", "201"):
            schema_range = find_schema_block_range(lines, mi, code)
            if schema_range is not None:
                schema_key_idx, end_excl = schema_range
                schema_key_indent = indent_level(lines[schema_key_idx])
                body_indent = schema_key_indent + 2
                new_body = schema_dict_to_body_lines(new_schema, body_indent)
                schema_patches.append(
                    SchemaPatch(
                        schema_key_idx=schema_key_idx,
                        end_excl=end_excl,
                        new_body=new_body,
                        marker_base_indent=schema_key_indent,
                    )
                )
                placed = True
                break
            inline_idx = find_inline_application_json_empty(lines, mi, code)
            if inline_idx is not None:
                ind = indent_level(lines[inline_idx])
                new_body = schema_dict_to_body_lines(new_schema, ind + 4)
                inline_patches.append(InlineJsonPatch(app_json_line_idx=inline_idx, new_body=new_body))
                placed = True
                break
        if not placed:
            continue

    if not schema_patches and not inline_patches:
        return list(lines), 0

    out = list(lines)
    work: list[tuple[int, str, SchemaPatch | InlineJsonPatch]] = []
    for p in schema_patches:
        work.append((p.schema_key_idx, "schema", p))
    for p in inline_patches:
        work.append((p.app_json_line_idx, "inline", p))

    for primary_idx, kind, p in sorted(work, key=lambda t: -t[0]):
        if kind == "schema":
            assert isinstance(p, SchemaPatch)
            del out[p.schema_key_idx + 1 : p.end_excl]
            for offset, nl in enumerate(p.new_body):
                out.insert(p.schema_key_idx + 1 + offset, nl)

            insert_at = p.schema_key_idx + 1 + len(p.new_body)
            marker_line = " " * p.marker_base_indent + f"{marker_key}: true"
            window = "\n".join(out[max(0, insert_at - 3) : insert_at + 8])
            if marker_key not in window:
                out.insert(insert_at, marker_line)
        else:
            assert isinstance(p, InlineJsonPatch)
            idx = p.app_json_line_idx
            ind = indent_level(out[idx])
            inner = ind + 2
            block: list[str] = [
                " " * ind + "application/json:",
                " " * inner + "schema:",
            ]
            block.extend(p.new_body)
            block.append(" " * inner + f"{marker_key}: true")
            del out[idx]
            for offset, nl in enumerate(block):
                out.insert(idx + offset, nl)

    n = len(schema_patches) + len(inline_patches)
    return out, n


def main() -> int:
    root = Path(__file__).resolve().parents[1]
    jsonl = Path(sys.argv[1]) if len(sys.argv) > 1 else root / "test/integration/_live_shapes/run.jsonl"
    openapi_path = Path(sys.argv[2]) if len(sys.argv) > 2 else root / "openapi/openapi.yaml"

    if not jsonl.is_file():
        print(f"Missing NDJSON log: {jsonl}", file=sys.stderr)
        print("Run integration tests with XURRENT_STRUCTURE_LOG_FILE set.", file=sys.stderr)
        return 1

    rows = load_ndjson(jsonl)
    agg = aggregate_skeletons(rows)
    if not agg:
        print("No response skeletons found in log.", file=sys.stderr)
        return 1

    text = openapi_path.read_text(encoding="utf-8")
    lines = text.splitlines(keepends=False)
    new_lines, n = surgical_patch_openapi_text(lines, agg)
    if n == 0:
        print("No operations matched paths in the OpenAPI document.", file=sys.stderr)
        return 1

    openapi_path.write_text("\n".join(new_lines) + ("\n" if text.endswith("\n") else ""), encoding="utf-8")

    api_copy = root / "api/openapi.yaml"
    if api_copy.is_file():
        api_copy.write_text(openapi_path.read_text(encoding="utf-8"), encoding="utf-8")

    print(f"Surgically updated {n} operation(s); merged {len(agg)} path/method skeleton(s).")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
