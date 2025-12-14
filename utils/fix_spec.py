import yaml
import sys

INPUT_FILE = '../openapi/openapi.yaml'
OUTPUT_FILE = '../openapi/openapi.yaml'

def fix_spec():
    print(f"Reading {INPUT_FILE}...")
    with open(INPUT_FILE, 'r', encoding='utf-8') as f:
        spec = yaml.safe_load(f)

    # Fix 1 & 2: requestBody content with no media type
    # We need to find paths with requestBody that have empty content or missing content
    if 'paths' in spec:
        for path, methods in spec['paths'].items():
            for method, operation in methods.items():
                if 'requestBody' in operation:
                    rb = operation['requestBody']
                    if 'content' in rb and not rb['content']:
                        # If content is empty, remove requestBody or add a default media type
                        # Assuming application/json for now
                        print(f"Fixing empty content in {path} {method}")
                        rb['content'] = {'application/json': {'schema': {'type': 'object'}}}
                    elif 'content' not in rb:
                         print(f"Fixing missing content in {path} {method}")
                         rb['content'] = {'application/json': {'schema': {'type': 'object'}}}

    # Fix 3: responses.undefined.description is missing
    # This usually happens when there is a response key that is not a status code or 'default'
    # Or if a response is missing description
    if 'paths' in spec:
        for path, methods in spec['paths'].items():
            for method, operation in methods.items():
                if 'responses' in operation:
                    # Check for 'undefined' response code and rename to 'default'
                    if 'undefined' in operation['responses']:
                        print(f"Fixing undefined response code in {path} {method}")
                        operation['responses']['default'] = operation['responses'].pop('undefined')
                    
                    for status, response in operation['responses'].items():
                        if 'description' not in response:
                            print(f"Fixing missing description in {path} {method} response {status}")
                            response['description'] = 'No description provided'

    print(f"Saving to {OUTPUT_FILE}...")
    with open(OUTPUT_FILE, 'w', encoding='utf-8') as f:
        yaml.dump(spec, f, default_flow_style=False, sort_keys=False)

if __name__ == '__main__':
    fix_spec()