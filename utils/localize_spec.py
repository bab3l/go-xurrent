import yaml
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

def is_ascii(s):
    return all(ord(c) < 128 for c in s)

def localize_value(key, value):
    if isinstance(value, str):
        if not is_ascii(value):
            # Create a generic English value based on the key if possible
            if key:
                return f"Example {key.capitalize()}"
            return "Example Value"
    return value

def traverse_and_localize(data, parent_key=None):
    if isinstance(data, dict):
        for key, value in data.items():
            if isinstance(value, (dict, list)):
                traverse_and_localize(value, key)
            else:
                data[key] = localize_value(key, value)
    elif isinstance(data, list):
        for i, item in enumerate(data):
            if isinstance(item, (dict, list)):
                traverse_and_localize(item, parent_key)
            else:
                data[i] = localize_value(parent_key, item)

def localize_spec(input_file, output_file):
    print(f"Reading {input_file}...")
    with open(input_file, 'r', encoding='utf-8') as f:
        data = yaml.load(f, Loader=yaml.Loader)

    print("Localizing content...")
    traverse_and_localize(data)

    print(f"Saving to {output_file}...")
    with open(output_file, 'w', encoding='utf-8') as f:
        yaml.dump(data, f, Dumper=yaml.Dumper, allow_unicode=True)
    
    print("Done.")

if __name__ == "__main__":
    localize_spec("openapi/openapi.yaml", "openapi/openapi_localized.yaml")
