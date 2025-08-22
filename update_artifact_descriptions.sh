#!/bin/bash
# Carefully update artifact descriptions for CyFir

echo "=== Updating Artifact Descriptions ==="
echo "Only updating user-visible descriptions, not technical content"
echo

UPDATED=0

# Find and update artifact files
find ./artifacts/definitions -name "*.yaml" -type f | while read artifact_file; do
    if grep -q "Velociraptor" "$artifact_file" 2>/dev/null; then
        # Create backup
        cp "$artifact_file" "$artifact_file.bak"
        
        # Update descriptions only (safe changes)
        sed -i 's/Since Velociraptor typically/Since CyFir typically/g' "$artifact_file"
        sed -i 's/Velociraptor server/CyFir server/g' "$artifact_file"
        sed -i 's/Velociraptor client/CyFir client/g' "$artifact_file"
        sed -i 's/Velociraptor agent/CyFir agent/g' "$artifact_file"
        sed -i 's/Velociraptor binary/CyFir binary/g' "$artifact_file"
        sed -i 's/Velociraptor installation/CyFir installation/g' "$artifact_file"
        sed -i 's/Velociraptor instance/CyFir instance/g' "$artifact_file"
        sed -i 's/Velociraptor deployment/CyFir deployment/g' "$artifact_file"
        sed -i 's/run Velociraptor/run CyFir/g' "$artifact_file"
        sed -i 's/running Velociraptor/running CyFir/g' "$artifact_file"
        
        # Check if file changed
        if ! diff -q "$artifact_file" "$artifact_file.bak" > /dev/null; then
            echo "Updated: $artifact_file"
            ((UPDATED++))
            rm "$artifact_file.bak"
        else
            mv "$artifact_file.bak" "$artifact_file"
        fi
    fi
done

echo
echo "=== Artifact Description Updates Complete ==="
echo "Files updated: $UPDATED"