#!/usr/bin/env python3

import yaml
import os

SPEC_PATH = '../openapi/openapi.yaml'

if not os.path.exists(SPEC_PATH):
    print(f"Spec file not found at {SPEC_PATH}")
    exit(1)

print(f"Patching {SPEC_PATH}...")

# Load the spec file
with open(SPEC_PATH, 'r', encoding='utf-8') as file:
    data = yaml.load(file, Loader=yaml.CLoader)

# TODO: Add patching logic here
# Example:
# if 'components' in data and 'schemas' in data['components']:
#     for schema in data['components']['schemas'].values():
#         # Fix something
#         pass

# Save the spec file
with open(SPEC_PATH, 'w', encoding='utf-8') as file:
    yaml.dump(data, file, Dumper=yaml.CDumper, sort_keys=False)

print("Spec patched successfully.")
