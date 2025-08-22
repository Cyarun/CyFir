#!/bin/bash

# Update documentation URLs to point to CyFir documentation
echo "Updating documentation URLs..."

# Update in Go files
find . -name "*.go" -type f | while read -r file; do
    if [[ "$file" == *"/vendor/"* ]]; then
        continue
    fi
    sed -i 's|https://docs\.velociraptor\.app|https://docs.cyfir.io|g' "$file"
    sed -i 's|docs\.velociraptor\.velocidex\.com|docs.cyfir.io|g' "$file"
done

# Update in YAML files
find . -name "*.yaml" -type f | while read -r file; do
    sed -i 's|https://docs\.velociraptor\.app|https://docs.cyfir.io|g' "$file"
    sed -i 's|docs\.velociraptor\.velocidex\.com|docs.cyfir.io|g' "$file"
done

# Update in MD files
find . -name "*.md" -type f | while read -r file; do
    sed -i 's|https://docs\.velociraptor\.app|https://docs.cyfir.io|g' "$file"
    sed -i 's|docs\.velociraptor\.velocidex\.com|docs.cyfir.io|g' "$file"
done

echo "Documentation URLs updated successfully!"