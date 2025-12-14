$ErrorActionPreference = "Stop"

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProjectRoot = Split-Path -Parent $ScriptDir
$OpenApiDir = Join-Path $ProjectRoot "openapi"
$InputFile = Join-Path $OpenApiDir "postman_collection.json"
$OutputFile = Join-Path $OpenApiDir "openapi.yaml"

Write-Host "Checking for Postman collection at $InputFile..."
if (-not (Test-Path $InputFile)) {
    Write-Error "Postman collection not found! Please place the JSON export at $InputFile"
}

Write-Host "Converting Postman collection to OpenAPI..."
# Use local npx since npm is installed
try {
    # Check if postman-to-openapi is installed, if not npx will ask to install it. 
    # We use --yes to auto-confirm installation of the package.
    npx --yes postman-to-openapi "$InputFile" "$OutputFile"
}
catch {
    Write-Error "Failed to run postman-to-openapi: $_"
}

if (Test-Path $OutputFile) {
    Write-Host "Successfully generated $OutputFile"
} else {
    Write-Error "Failed to generate OpenAPI spec"
}
