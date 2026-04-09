$ErrorActionPreference = "Stop"

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProjectRoot = Split-Path -Parent $ScriptDir
$OpenApiDir = Join-Path $ProjectRoot "openapi"
$InputFile = Join-Path $OpenApiDir "openapi.yaml"
$OutputDir = $ProjectRoot

Write-Host "Generating Go client from $InputFile..."

# Pin image tag for reproducible builds (see .openapi-generator/VERSION for last generated toolchain).
$OpenApiGenImage = "openapitools/openapi-generator-cli:v7.8.0"

# Run OpenAPI Generator via Docker
docker run --rm `
    -v "${ProjectRoot}:/local" `
    "${OpenApiGenImage}" generate `
    -i /local/openapi/openapi.yaml `
    -g go `
    -o /local `
    --git-user-id "xurrent" `
    --git-repo-id "go-xurrent" `
    --package-name "xurrent" `
    -p packageName=xurrent `
    -p isGoSubmodule=false `
    -p enumClassPrefix=true `
    -p useOneOfDiscriminatorLookup=true

Write-Host "Client generation complete."