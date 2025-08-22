#!/bin/bash
# Carefully update Go file comments for CyFir

echo "=== Updating Go File Comments ==="
echo "This script only updates comments, not code"
echo

# Counter
UPDATED=0

# Update single-line comments
echo "Updating single-line comments..."
find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./node_modules/*" | while read file; do
    if grep -q "// .*Velociraptor" "$file" 2>/dev/null; then
        # Create a backup
        cp "$file" "$file.bak"
        
        # Update comments but not URLs or code
        sed -i 's|// \(.*\)Velociraptor server|// \1CyFir server|g' "$file"
        sed -i 's|// \(.*\)Velociraptor client|// \1CyFir client|g' "$file"
        sed -i 's|// \(.*\)Velociraptor CA|// \1CyFir CA|g' "$file"
        sed -i 's|// \(.*\)Velociraptor service|// \1CyFir service|g' "$file"
        sed -i 's|// \(.*\)Velociraptor needs|// \1CyFir needs|g' "$file"
        sed -i 's|// \(.*\)Velociraptor allows|// \1CyFir allows|g' "$file"
        sed -i 's|// \(.*\)Velociraptor has|// \1CyFir has|g' "$file"
        sed -i 's|// \(.*\)Velociraptor is|// \1CyFir is|g' "$file"
        sed -i 's|// \(.*\)Velociraptor uses|// \1CyFir uses|g' "$file"
        sed -i 's|// \(.*\)Velociraptor will|// \1CyFir will|g' "$file"
        sed -i 's|// \(.*\)Velociraptor can|// \1CyFir can|g' "$file"
        sed -i 's|// \(.*\)Velociraptor does|// \1CyFir does|g' "$file"
        sed -i 's|// \(.*\)Velociraptor should|// \1CyFir should|g' "$file"
        sed -i 's|// \(.*\)Velociraptor only|// \1CyFir only|g' "$file"
        
        # Check if file actually changed
        if ! diff -q "$file" "$file.bak" > /dev/null; then
            echo "Updated: $file"
            ((UPDATED++))
            rm "$file.bak"
        else
            # No changes, restore original
            mv "$file.bak" "$file"
        fi
    fi
done

echo
echo "=== Go Comment Updates Complete ==="
echo "Files updated: $UPDATED"
echo "Note: Only comments were changed, not code or function names"