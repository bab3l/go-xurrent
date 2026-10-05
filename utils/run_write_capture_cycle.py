#!/usr/bin/env python3
"""
Write-verb NDJSON capture: conservative rate limits + live capture pipeline (integration + GET sweep + optional merge).

Forces defaults only when unset (override via .env or shell):
  XURRENT_HTTP_MIN_INTERVAL — Go integration client spacing + 429 handling (config_helpers.go)
  XURRENT_SWEEP_DELAY       — GET sweep spacing (collect_live_response_shapes.py)

Forwards all CLI arguments to utils/live_capture_pipeline.py.

Typical full run (review tenant + token before executing):
  python utils/run_write_capture_cycle.py --fresh -v

Targeted writes only (no GET sweep / merge):
  python utils/run_write_capture_cycle.py --integration-profile capture-writes \\
    --go-test-run "TestWriteCapture_|TestWriteShape_|TestMutation_" -v \\
    --skip-sweep --skip-merge --skip-report

Optional PostImport (tenant-specific CSV): set XURRENT_WRITE_CAPTURE_POST_IMPORT=1 and paths in .env.
"""

from __future__ import annotations

import os
import subprocess
import sys
from pathlib import Path


def load_dotenv(path: Path) -> None:
    if not path.is_file():
        return
    for raw in path.read_text(encoding="utf-8").splitlines():
        line = raw.strip()
        if not line or line.startswith("#") or "=" not in line:
            continue
        key, _, val = line.partition("=")
        key, val = key.strip(), val.strip().strip('"').strip("'")
        if key and key not in os.environ:
            os.environ[key] = val


def main() -> int:
    root = Path(__file__).resolve().parents[1]
    load_dotenv(root / ".env")
    os.environ.setdefault("XURRENT_HTTP_MIN_INTERVAL", "0.4")
    os.environ.setdefault("XURRENT_SWEEP_DELAY", "0.35")
    pipeline = root / "utils" / "live_capture_pipeline.py"
    cmd = [sys.executable, str(pipeline), *sys.argv[1:]]
    print("+", " ".join(cmd), file=sys.stderr, flush=True)
    return subprocess.call(cmd, cwd=root)


if __name__ == "__main__":
    raise SystemExit(main())
