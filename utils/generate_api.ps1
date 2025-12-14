$ErrorActionPreference = "Stop"

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProjectRoot = Split-Path -Parent $ScriptDir

Write-Host "Step 1: Fixing OpenAPI spec..."
# Run from ProjectRoot so relative paths in python script work
Push-Location $ProjectRoot
python utils/fix_spec.py
if ($LASTEXITCODE -eq 0) {
    Move-Item -Path "openapi/openapi_fixed.yaml" -Destination "openapi/openapi.yaml" -Force
} else {
    Write-Error "Fix spec failed"
}
Pop-Location

Write-Host "Step 1.5: Cleaning up stale files..."
if (Test-Path "$ProjectRoot/docs/Class*.md") {
    Remove-Item "$ProjectRoot/docs/Class*.md" -Force
}
if (Test-Path "$ProjectRoot/test/api_class*.go") {
    Remove-Item "$ProjectRoot/test/api_class*.go" -Force
}

Write-Host "Step 2: Generating Go client..."
& "$ScriptDir\generate_client.ps1"

Write-Host "Step 3: Tidy and Format..."
Push-Location $ProjectRoot
go mod tidy
go fmt ./...
Pop-Location

Write-Host "API generation complete!"
