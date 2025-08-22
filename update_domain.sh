#!/bin/bash

# Update domain references from docs.cyfir.io to cyfir.cynorsense.com

echo "Updating domain references to cyfir.cynorsense.com..."

# Find and replace in all Go files
find . -name "*.go" -type f -exec sed -i 's/docs\.cyfir\.io/cyfir.cynorsense.com/g' {} \;

# Find and replace in all YAML files  
find . -name "*.yaml" -type f -exec sed -i 's/docs\.cyfir\.io/cyfir.cynorsense.com/g' {} \;
find . -name "*.yml" -type f -exec sed -i 's/docs\.cyfir\.io/cyfir.cynorsense.com/g' {} \;

# Find and replace in all Markdown files
find . -name "*.md" -type f -exec sed -i 's/docs\.cyfir\.io/cyfir.cynorsense.com/g' {} \;

# Update any references to www.cyfir.io to cynorsense.com
find . -name "*.go" -type f -exec sed -i 's/www\.cyfir\.io/cynorsense.com/g' {} \;
find . -name "*.yaml" -type f -exec sed -i 's/www\.cyfir\.io/cynorsense.com/g' {} \;
find . -name "*.yml" -type f -exec sed -i 's/www\.cyfir\.io/cynorsense.com/g' {} \;
find . -name "*.md" -type f -exec sed -i 's/www\.cyfir\.io/cynorsense.com/g' {} \;

echo "Domain update complete!"