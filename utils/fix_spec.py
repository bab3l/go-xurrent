import yaml
import json
import re
import sys
from collections import OrderedDict

# Configure yaml to preserve order
def setup_yaml():
    def dict_representer(dumper, data):
        return dumper.represent_dict(data.items())
    def dict_constructor(loader, node):
        return OrderedDict(loader.construct_pairs(node))
    
    yaml.add_representer(OrderedDict, dict_representer)
    yaml.add_constructor(yaml.resolver.BaseResolver.DEFAULT_MAPPING_TAG, dict_constructor)

setup_yaml()

def load_mapping(mapping_file):
    with open(mapping_file, 'r') as f:
        return json.load(f)

def normalize_path(path):
    # Remove trailing slash
    if path.endswith('/') and len(path) > 1:
        path = path[:-1]
    
    segments = path.split('/')
    new_segments = []
    for segment in segments:
        # Check for numeric ID
        if re.match(r'^\d+$', segment):
            new_segments.append('{id}')
        # Check for UUID (simple check)
        elif re.match(r'^[a-f0-9]{32}$', segment):
            new_segments.append('{token}')
        # Check for query param in path (e.g. manager=53)
        elif '=' in segment:
            # This is likely a demo path that shouldn't exist in the spec as a path
            # We'll mark it for removal or manual check
            return None
        else:
            new_segments.append(segment)
    
    return '/'.join(new_segments)

def fix_spec(input_file, output_file, mapping_file):
    tag_mapping = load_mapping(mapping_file)
    
    with open(input_file, 'r', encoding='utf-8') as f:
        data = yaml.load(f, Loader=yaml.Loader)

    # 1. Fix Tags
    if 'tags' in data:
        new_tags = []
        seen_tags = set()
        for tag in data['tags']:
            old_name = tag['name']
            new_name = tag_mapping.get(old_name, old_name)
            if new_name not in seen_tags:
                tag['name'] = new_name
                new_tags.append(tag)
                seen_tags.add(new_name)
        data['tags'] = new_tags

    # 2. Fix Paths (Normalize and Deduplicate)
    new_paths = OrderedDict()
    
    for path, methods in data.get('paths', {}).items():
        normalized_path = normalize_path(path)
        
        if normalized_path is None:
            print(f"Skipping invalid path: {path}")
            continue
            
        if normalized_path not in new_paths:
            new_paths[normalized_path] = methods
        else:
            print(f"Merging duplicate path: {path} -> {normalized_path}")
            # Merge methods. If method exists, keep existing (or overwrite? usually first is better in generated specs)
            for method, details in methods.items():
                if method not in new_paths[normalized_path]:
                    new_paths[normalized_path][method] = details
                else:
                    print(f"  Duplicate method {method} for {normalized_path}, skipping.")

        # Update tags in methods
        for method, details in new_paths[normalized_path].items():
            if 'tags' in details:
                details['tags'] = [tag_mapping.get(t, t) for t in details['tags']]
            
            # Fix parameters if we introduced {id} or {token}
            if '{id}' in normalized_path or '{token}' in normalized_path:
                if 'parameters' not in details:
                    details['parameters'] = []
                
                # Check if id/token param exists
                param_names = [p.get('name') for p in details['parameters']]
                
                if '{id}' in normalized_path and 'id' not in param_names:
                    details['parameters'].insert(0, {
                        'name': 'id',
                        'in': 'path',
                        'required': True,
                        'schema': {'type': 'integer'}
                    })
                
                if '{token}' in normalized_path and 'token' not in param_names:
                    details['parameters'].insert(0, {
                        'name': 'token',
                        'in': 'path',
                        'required': True,
                        'schema': {'type': 'string'}
                    })

    data['paths'] = new_paths

    with open(output_file, 'w', encoding='utf-8') as f:
        yaml.dump(data, f, Dumper=yaml.Dumper, allow_unicode=True)
    
    print(f"Saved fixed spec to {output_file}")

if __name__ == "__main__":
    fix_spec("openapi/openapi.yaml", "openapi/openapi_fixed.yaml", "utils/tag_mapping.json")
