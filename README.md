# go-xurrent

Go client library for the Xurrent (formerly 4me) ITSM platform.

## Development Process

This library is generated from an OpenAPI specification derived from a Postman collection.

### Prerequisites

- Docker (for running `openapi-generator`)
- Node.js & npm (for running `postman-to-openapi`)
- Go 1.24+
- Python 3 (for spec patching)

### Workflow

1.  **Prepare OpenAPI Spec**:
    - Place your OpenAPI specification file (e.g., `openapi-v1.0.0.yaml`) in the `openapi/` directory.
    - The generation script will automatically pick the latest versioned file.

2.  **Generate Client**:
    - Run `utils/generate_api.ps1`.
    - This script will:
        - Select the latest `openapi-*.yaml` file.
        - Run `utils/fix-spec.py` to patch the spec.
        - Run the OpenAPI Generator to create the Go client.

### Directory Structure

- `openapi/`: Contains the input OpenAPI specification files.
- `utils/`: Scripts for patching and generation.
- `patches/`: Manual Go code patches to apply after generation.
