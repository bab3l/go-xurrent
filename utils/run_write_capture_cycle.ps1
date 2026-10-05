# Write-verb capture wrapper — see utils/run_write_capture_cycle.py
$ErrorActionPreference = "Stop"
$here = Split-Path -Parent $MyInvocation.MyCommand.Path
$root = Resolve-Path (Join-Path $here "..")
Set-Location $root
python (Join-Path $here "run_write_capture_cycle.py") @args
exit $LASTEXITCODE
