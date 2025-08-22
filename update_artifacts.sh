#!/bin/bash
# CyFir Artifact Update Script
# This script helps update artifact definitions with CyFir branding

echo "=== CyFir Artifact Update Tool ==="
echo

# Counter for tracking progress
UPDATED=0
TOTAL=0

# Function to update a single artifact file
update_artifact() {
    local file=$1
    local changed=false
    
    # Check if file contains Velociraptor references
    if grep -q "Velociraptor" "$file"; then
        echo "Updating: $file"
        
        # Update common references
        sed -i 's/Velociraptor server/CyFir server/g' "$file"
        sed -i 's/Velociraptor client/CyFir client/g' "$file"
        sed -i 's/Velociraptor deployment/CyFir deployment/g' "$file"
        sed -i 's/Velociraptor GUI/CyFir GUI/g' "$file"
        sed -i 's/Velociraptor instance/CyFir instance/g' "$file"
        sed -i 's/Velociraptor Query Language/Velociraptor Query Language/g' "$file" # Keep VQL name
        sed -i 's/the Velociraptor/the CyFir/g' "$file"
        sed -i 's/The Velociraptor/The CyFir/g' "$file"
        sed -i 's/with Velociraptor/with CyFir/g' "$file"
        sed -i 's/by Velociraptor/by CyFir/g' "$file"
        sed -i 's/from Velociraptor/from CyFir/g' "$file"
        sed -i 's/in Velociraptor/in CyFir/g' "$file"
        sed -i 's/using Velociraptor/using CyFir/g' "$file"
        
        # Update URLs
        sed -i 's|docs.velociraptor.app|cyfir.cynorsense.com/docs|g' "$file"
        sed -i 's|github.com/Velocidex/velociraptor|github.com/Cyarun/CyFir|g' "$file"
        
        changed=true
        ((UPDATED++))
    fi
    
    ((TOTAL++))
}

# Process artifacts based on priority
echo "Processing high-priority server artifacts..."
for file in artifacts/definitions/Server/*.yaml; do
    [ -f "$file" ] && update_artifact "$file"
done

echo
echo "Processing client management artifacts..."
for file in artifacts/definitions/Admin/Client/*.yaml; do
    [ -f "$file" ] && update_artifact "$file"
done

echo
echo "Processing generic client artifacts..."
for file in artifacts/definitions/Generic/Client/*.yaml; do
    [ -f "$file" ] && update_artifact "$file"
done

echo
echo "Processing Windows system artifacts..."
for file in artifacts/definitions/Windows/System/*.yaml; do
    [ -f "$file" ] && update_artifact "$file"
done

echo
echo "Processing detection artifacts..."
for file in artifacts/definitions/Windows/Detection/*.yaml artifacts/definitions/Generic/Detection/*.yaml; do
    [ -f "$file" ] && update_artifact "$file"
done

echo
echo "=== Summary ==="
echo "Total artifacts processed: $TOTAL"
echo "Artifacts updated: $UPDATED"
echo "Completion: $((UPDATED * 100 / TOTAL))%"
echo
echo "Run 'git diff' to review changes"
echo "Run 'git add -A && git commit' to commit changes"