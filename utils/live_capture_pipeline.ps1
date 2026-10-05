# Live NDJSON capture pipeline — see utils/live_capture_pipeline.py
# Usage (repo root):  .\utils\live_capture_pipeline.ps1
#                      .\utils\live_capture_pipeline.ps1 --fresh
$ErrorActionPreference = "Stop"
$here = Split-Path -Parent $MyInvocation.MyCommand.Path
$root = Resolve-Path (Join-Path $here "..")
Set-Location $root
python (Join-Path $here "live_capture_pipeline.py") @args
exit $LASTEXITCODE
