$ErrorActionPreference = "Stop"

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProjectRoot = Split-Path -Parent $ScriptDir
$OpenApiDir = Join-Path $ProjectRoot "openapi"
$SpecFile = Join-Path $OpenApiDir "openapi.yaml"

# 1. Find and prepare OpenAPI spec
Write-Host "Preparing OpenAPI spec..."

# Find the OpenAPI spec file
$specFiles = Get-ChildItem "$OpenApiDir" -Filter "openapi-*.yaml" | Sort-Object Name -Descending
if ($specFiles.Count -eq 0) {
    # Fallback to openapi.yaml if no versioned file found
    if (Test-Path "$OpenApiDir\openapi.yaml") {
        Write-Host "Using default openapi.yaml"
    } else {
        Write-Error "No OpenAPI spec found in $OpenApiDir. Please provide an openapi-*.yaml or openapi.yaml file."
        exit 1
    }
} else {
    $latestSpec = $specFiles[0]
    Write-Host "Using OpenAPI spec: $($latestSpec.Name)"
    Copy-Item $latestSpec.FullName "$OpenApiDir\openapi.yaml" -Force
}

# 2. Patch the OpenAPI spec
Write-Host "Patching OpenAPI spec..."
Push-Location $ScriptDir
python fix-spec.py
Pop-Location

# 3. Generate Go Client
Write-Host "Generating Go client..."
# Clean up old files
if (Test-Path "$ProjectRoot\api") {
    Remove-Item "$ProjectRoot\api" -Recurse -Force
}
if (Test-Path "$ProjectRoot\docs") {
    Remove-Item "$ProjectRoot\docs" -Recurse -Force
}

# Run OpenAPI Generator
docker run --rm `
    -v "${ProjectRoot}:/local" `
    openapitools/openapi-generator-cli:v7.11.0 `
    generate `
    --input-spec /local/openapi/openapi.yaml `
    --generator-name go `
    --output /local `
    --package-name xurrent `
    --git-user-id "your-org" `
    --git-repo-id "go-xurrent" `
    --additional-properties=isGoSubmodule=true,packageName=xurrent

# 4. Tidy and Format
Write-Host "Running go mod tidy..."
Push-Location $ProjectRoot
go mod tidy
go fmt ./...
Pop-Location

Write-Host "Generation complete!"
