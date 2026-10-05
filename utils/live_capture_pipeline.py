#!/usr/bin/env python3
"""
Live capture pipeline: integration tests with structure logging, GET sweep, merge into OpenAPI,
YAML validation, and refresh of the NDJSON gap report.

Loads repo-root .env into the environment (does not override variables already set).

Usage (repo root):
  python utils/live_capture_pipeline.py
  python utils/live_capture_pipeline.py --fresh              # truncate NDJSON before tests
  python utils/live_capture_pipeline.py --skip-integration   # sweep + merge only
  python utils/live_capture_pipeline.py --skip-merge         # tests + sweep, no OpenAPI edit

Targeted integration (faster iteration):
  python utils/live_capture_pipeline.py --go-test-run "TestWriteShape_SearchGET" \\
      --integration-profile capture-writes --skip-sweep --skip-merge --skip-report -v

  python utils/live_capture_pipeline.py --parent-fixture --go-test-run TestFixture_ParentIDsForCapture \\
      --skip-sweep --skip-merge --skip-report -v

  python utils/live_capture_pipeline.py --go-test-run "TestFixture_ParentIDsForCapture|TestWriteShape_" -v

Requires: Go toolchain, pyyaml, ruamel.yaml (for merge), XURRENT_TOKEN + XURRENT_ACCOUNT in .env
          for integration tests and sweep (unless those steps are skipped).

Integration profiles:
  full            — mutations, lifecycle, endpoint coverage, write-shape (default; longest)
  capture-writes  — mutations + write-shape + structure log only (skip lifecycle + broad GET coverage)

Rate limits (defaults applied when unset, same as run_write_capture_cycle.py):
  XURRENT_HTTP_MIN_INTERVAL — Go client spacing + 429 handling (default 0.4s).
  XURRENT_SWEEP_DELAY       — delay between GET sweep requests (default 0.35s).

Optional cooldown before integration or sweep:
  XURRENT_LIVE_CAPTURE_PRE_SLEEP — seconds after loading .env (e.g. after a heavy run). The legacy name
  XURRENT_PHASE1_PRE_SLEEP is still accepted.

For a one-liner that only sets rate env defaults, utils/run_write_capture_cycle.py forwards to this script.

Integration tests set: STRUCTURE_LOG, ALLOW_MUTATIONS, WRITE_SHAPE_COVERAGE, and (full only)
LIFECYCLE_MUTATIONS + ENDPOINT_COVERAGE.
"""

from __future__ import annotations

import argparse
import os
import subprocess
import sys
import time
from pathlib import Path

_UTILS_DIR = Path(__file__).resolve().parent
if str(_UTILS_DIR) not in sys.path:
    sys.path.insert(0, str(_UTILS_DIR))
from doc_paths import NDJSON_GAP_REPORT_MD  # noqa: E402


def apply_rate_limit_defaults() -> None:
    """Match utils/run_write_capture_cycle.py: conservative spacing when not set in .env or shell."""
    os.environ.setdefault("XURRENT_HTTP_MIN_INTERVAL", "0.4")
    os.environ.setdefault("XURRENT_SWEEP_DELAY", "0.35")


def resolve_pre_run_sleep_seconds(cli_value: float | None) -> float:
    if cli_value is not None:
        return max(0.0, float(cli_value))
    for key in ("XURRENT_LIVE_CAPTURE_PRE_SLEEP", "XURRENT_PHASE1_PRE_SLEEP"):
        raw = os.environ.get(key, "").strip()
        if raw:
            try:
                return max(0.0, float(raw))
            except ValueError:
                print(
                    f"Ignoring invalid {key}={raw!r} (expected a number of seconds).",
                    file=sys.stderr,
                )
            break
    return 0.0


def load_dotenv(path: Path) -> None:
    if not path.is_file():
        return
    for raw in path.read_text(encoding="utf-8").splitlines():
        line = raw.strip()
        if not line or line.startswith("#"):
            continue
        if "=" not in line:
            continue
        key, _, val = line.partition("=")
        key, val = key.strip(), val.strip().strip('"').strip("'")
        if key and key not in os.environ:
            os.environ[key] = val


def load_env_file_override(path: Path) -> None:
    """Load KEY=value into os.environ (override existing) for fixture files."""
    if not path.is_file():
        return
    for raw in path.read_text(encoding="utf-8").splitlines():
        line = raw.strip()
        if not line or line.startswith("#"):
            continue
        if "=" not in line:
            continue
        key, _, val = line.partition("=")
        key, val = key.strip(), val.strip().strip('"').strip("'")
        if key:
            os.environ[key] = val


def run(cmd: list[str], cwd: Path, **kwargs) -> None:
    print("+", " ".join(cmd), file=sys.stderr, flush=True)
    subprocess.run(cmd, cwd=cwd, check=True, **kwargs)


def main() -> int:
    root = Path(__file__).resolve().parents[1]
    ap = argparse.ArgumentParser(description=__doc__)
    ap.add_argument(
        "--env-file",
        type=Path,
        default=root / ".env",
        help="Dotenv file to load if present (default: .env)",
    )
    ap.add_argument(
        "--ndjson",
        type=Path,
        default=root / "test/integration/_live_shapes/full_run.jsonl",
        help="NDJSON path for structure log + sweep output",
    )
    ap.add_argument(
        "--fresh",
        action="store_true",
        help="Truncate --ndjson before integration tests (clean run)",
    )
    ap.add_argument("--skip-integration", action="store_true", help="Skip go test (sweep + merge only)")
    ap.add_argument("--skip-sweep", action="store_true", help="Skip collect_live_response_shapes.py")
    ap.add_argument("--skip-merge", action="store_true", help="Skip merge into openapi + api copy + yaml check")
    ap.add_argument(
        "--skip-report",
        action="store_true",
        help=f"Skip report_ndjson_openapi_gaps.py --write ({NDJSON_GAP_REPORT_MD.relative_to(root)})",
    )
    ap.add_argument(
        "--no-seed-sweep",
        action="store_true",
        help="Pass --no-seed to collect_live_response_shapes.py (no POST bootstrap)",
    )
    ap.add_argument(
        "-v",
        "--verbose",
        action="store_true",
        help="Pass -v to go test (more logging)",
    )
    ap.add_argument(
        "--go-test-run",
        dest="go_test_run",
        default=None,
        metavar="REGEX",
        help="Passed to go test -run REGEX (e.g. TestWriteShape_|TestFixture_). Default: all tests in package.",
    )
    ap.add_argument(
        "--integration-profile",
        choices=("full", "capture-writes"),
        default="full",
        help="full = all capture env flags (default). capture-writes = mutations + write-shape only (faster).",
    )
    ap.add_argument(
        "--parent-fixture",
        action="store_true",
        help="Set XURRENT_BUILD_PARENT_FIXTURE=1 so TestFixture_ParentIDsForCapture can write parent_fixture.env",
    )
    ap.add_argument(
        "--load-parent-fixture",
        action="store_true",
        help="Merge test/integration/_live_shapes/parent_fixture.env into environment before go test",
    )
    ap.add_argument(
        "--pre-run-sleep",
        type=float,
        default=None,
        metavar="SECONDS",
        help="Sleep before integration/sweep (rate-limit cooldown). Overrides XURRENT_LIVE_CAPTURE_PRE_SLEEP.",
    )
    args = ap.parse_args()

    print(
        "Live capture pipeline: integration tests + GET sweep + merge (several minutes; sweep uses XURRENT_SWEEP_DELAY).",
        file=sys.stderr,
        flush=True,
    )

    load_dotenv(args.env_file if args.env_file.is_absolute() else root / args.env_file)
    apply_rate_limit_defaults()

    pre_sleep = resolve_pre_run_sleep_seconds(args.pre_run_sleep)
    if pre_sleep > 0:
        print(
            f"Pre-run backoff sleep: {pre_sleep}s (--pre-run-sleep or XURRENT_LIVE_CAPTURE_PRE_SLEEP)",
            file=sys.stderr,
            flush=True,
        )
        time.sleep(pre_sleep)

    min_iv = os.environ.get("XURRENT_HTTP_MIN_INTERVAL", "")
    sweep_d = os.environ.get("XURRENT_SWEEP_DELAY", "")
    print(
        f"Effective rate limits: XURRENT_HTTP_MIN_INTERVAL={min_iv!r} XURRENT_SWEEP_DELAY={sweep_d!r}",
        file=sys.stderr,
        flush=True,
    )

    ndjson = args.ndjson if args.ndjson.is_absolute() else root / args.ndjson
    ndjson.parent.mkdir(parents=True, exist_ok=True)

    fixture_path = root / "test/integration/_live_shapes/parent_fixture.env"
    if args.load_parent_fixture:
        load_env_file_override(fixture_path)
        if fixture_path.is_file():
            print(f"Loaded {fixture_path.relative_to(root)} into environment", file=sys.stderr)

    if args.fresh and ndjson.is_file():
        ndjson.unlink()
        print(f"Removed {ndjson.relative_to(root)} (--fresh)", file=sys.stderr)

    if not args.skip_integration:
        token = os.environ.get("XURRENT_TOKEN", "").strip()
        account = os.environ.get("XURRENT_ACCOUNT", "").strip()
        if not token or not account:
            print(
                "XURRENT_TOKEN and XURRENT_ACCOUNT must be set (e.g. via .env) for integration tests.",
                file=sys.stderr,
            )
            return 1

        os.environ["XURRENT_STRUCTURE_LOG_FILE"] = str(ndjson.resolve())
        os.environ["XURRENT_ALLOW_MUTATIONS"] = "1"
        os.environ["XURRENT_WRITE_SHAPE_COVERAGE"] = "1"
        if args.integration_profile == "full":
            os.environ["XURRENT_ALLOW_LIFECYCLE_MUTATIONS"] = "1"
            os.environ["XURRENT_ENDPOINT_COVERAGE"] = "1"
        else:
            os.environ.pop("XURRENT_ALLOW_LIFECYCLE_MUTATIONS", None)
            os.environ.pop("XURRENT_ENDPOINT_COVERAGE", None)

        if args.parent_fixture:
            os.environ["XURRENT_BUILD_PARENT_FIXTURE"] = "1"

        go_cmd = ["go", "test", "-tags=integration", "./test/integration/"]
        if args.go_test_run:
            go_cmd.extend(["-run", args.go_test_run])
        go_cmd.append("-count=1")
        if args.verbose:
            go_cmd.append("-v")

        run(go_cmd, cwd=root)

    if not args.skip_sweep:
        token = os.environ.get("XURRENT_TOKEN", "").strip()
        account = os.environ.get("XURRENT_ACCOUNT", "").strip()
        if not token or not account:
            print(
                "XURRENT_TOKEN and XURRENT_ACCOUNT required for GET sweep (or use --skip-sweep).",
                file=sys.stderr,
            )
            return 1
        collect_cmd = [
            sys.executable,
            str(root / "utils/collect_live_response_shapes.py"),
            "--out",
            str(ndjson),
        ]
        if args.no_seed_sweep:
            collect_cmd.append("--no-seed")
        run(collect_cmd, cwd=root)

    if not args.skip_merge:
        if not ndjson.is_file():
            print(f"Missing NDJSON {ndjson}; run integration and/or sweep first.", file=sys.stderr)
            return 1
        run(
            [
                sys.executable,
                str(root / "utils/merge_skeleton_into_openapi.py"),
                str(ndjson),
                str(root / "openapi/openapi.yaml"),
            ],
            cwd=root,
        )
        try:
            import yaml
        except ImportError:
            print("pip install pyyaml", file=sys.stderr)
            return 1
        openapi_path = root / "openapi/openapi.yaml"
        yaml.safe_load(openapi_path.read_text(encoding="utf-8"))
        print(f"YAML OK: {openapi_path.relative_to(root)}", file=sys.stderr)

    if not args.skip_report:
        run(
            [
                sys.executable,
                str(root / "utils/report_ndjson_openapi_gaps.py"),
                "--ndjson",
                str(ndjson),
                "--write",
                str(NDJSON_GAP_REPORT_MD),
            ],
            cwd=root,
        )

    print("Live capture pipeline finished.", file=sys.stderr)
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except subprocess.CalledProcessError as e:
        print(f"Command failed with exit code {e.returncode}", file=sys.stderr)
        raise SystemExit(e.returncode)
