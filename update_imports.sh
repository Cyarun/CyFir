#!/bin/bash

# Update all import paths from old module name to new module name
echo "Updating import paths from www.velocidex.com/golang/velociraptor to github.com/Cyarun/CyFir..."

# Find all Go files and update imports
find . -name "*.go" -type f | while read -r file; do
    # Skip vendor directory if it exists
    if [[ "$file" == *"/vendor/"* ]]; then
        continue
    fi
    
    # Update the import path
    sed -i 's|www\.velocidex\.com/golang/velociraptor|github.com/Cyarun/CyFir|g' "$file"
done

echo "Import paths updated successfully!"

# Count remaining references to verify
remaining=$(grep -r "www.velocidex.com/golang/velociraptor" --include="*.go" . | wc -l)
echo "Remaining references: $remaining"