$ErrorActionPreference = "Stop"

Write-Host "Step 1: Fixing OpenAPI spec..."
python fix_spec.py

Write-Host "Step 2: Generating Go client..."
.\generate_client.ps1

Write-Host "Step 3: Tidy and Format..."
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProjectRoot = Split-Path -Parent $ScriptDir
Push-Location $ProjectRoot
go mod tidy
go fmt ./...
Pop-Location

Write-Host "API generation complete!"
