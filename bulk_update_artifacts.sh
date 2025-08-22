#!/bin/bash
# CyFir Bulk Artifact Update Script
# Updates all remaining artifacts with CyFir branding

echo "=== CyFir Bulk Artifact Update ==="
echo "This will update ALL artifact files. Continue? (y/n)"
read -r response

if [[ ! "$response" =~ ^[Yy]$ ]]; then
    echo "Aborted."
    exit 0
fi

# Counter
UPDATED=0
SKIPPED=0
TOTAL=0

# Create backup
echo "Creating backup..."
tar -czf artifacts_backup_$(date +%Y%m%d_%H%M%S).tar.gz artifacts/

# Function to update artifact
update_artifact_file() {
    local file=$1
    local temp_file=$(mktemp)
    local changed=false
    
    # Skip already updated files
    if grep -q "CyFir" "$file" 2>/dev/null; then
        ((SKIPPED++))
        return
    fi
    
    # Perform replacements
    cp "$file" "$temp_file"
    
    # Context-aware replacements
    sed -i 's/Velociraptor server/CyFir server/g' "$temp_file"
    sed -i 's/Velociraptor client/CyFir client/g' "$temp_file"
    sed -i 's/Velociraptor agent/CyFir agent/g' "$temp_file"
    sed -i 's/Velociraptor binary/CyFir binary/g' "$temp_file"
    sed -i 's/Velociraptor deployment/CyFir deployment/g' "$temp_file"
    sed -i 's/Velociraptor GUI/CyFir GUI/g' "$temp_file"
    sed -i 's/Velociraptor instance/CyFir instance/g' "$temp_file"
    sed -i 's/Velociraptor installation/CyFir installation/g' "$temp_file"
    sed -i 's/Velociraptor configuration/CyFir configuration/g' "$temp_file"
    sed -i 's/Velociraptor service/CyFir service/g' "$temp_file"
    sed -i 's/Velociraptor process/CyFir process/g' "$temp_file"
    sed -i 's/Velociraptor endpoint/CyFir endpoint/g' "$temp_file"
    sed -i 's/Velociraptor artifacts/CyFir artifacts/g' "$temp_file"
    sed -i 's/Velociraptor API/CyFir API/g' "$temp_file"
    sed -i 's/Velociraptor frontend/CyFir frontend/g' "$temp_file"
    
    # Sentence replacements
    sed -i 's/the Velociraptor/the CyFir/g' "$temp_file"
    sed -i 's/The Velociraptor/The CyFir/g' "$temp_file"
    sed -i 's/with Velociraptor/with CyFir/g' "$temp_file"
    sed -i 's/by Velociraptor/by CyFir/g' "$temp_file"
    sed -i 's/from Velociraptor/from CyFir/g' "$temp_file"
    sed -i 's/in Velociraptor/in CyFir/g' "$temp_file"
    sed -i 's/using Velociraptor/using CyFir/g' "$temp_file"
    sed -i 's/for Velociraptor/for CyFir/g' "$temp_file"
    sed -i 's/of Velociraptor/of CyFir/g' "$temp_file"
    sed -i 's/to Velociraptor/to CyFir/g' "$temp_file"
    
    # Preserve VQL name
    sed -i 's/CyFir Query Language/Velociraptor Query Language/g' "$temp_file"
    
    # Update URLs
    sed -i 's|docs\.velociraptor\.app|cyfir.cynorsense.com/docs|g' "$temp_file"
    sed -i 's|github\.com/Velocidex/velociraptor|github.com/Cyarun/CyFir|g' "$temp_file"
    sed -i 's|www\.velocidex\.com|cynorsense.com|g' "$temp_file"
    
    # Update author if it's Velocidex
    sed -i 's/author: Velocidex/author: CynorSense Solutions/g' "$temp_file"
    sed -i 's/author: "Velocidex"/author: "CynorSense Solutions"/g' "$temp_file"
    
    # Check if file changed
    if ! cmp -s "$file" "$temp_file"; then
        mv "$temp_file" "$file"
        echo "Updated: $file"
        ((UPDATED++))
        changed=true
    else
        rm "$temp_file"
        ((SKIPPED++))
    fi
}

# Process all YAML files
echo
echo "Processing artifact files..."
while IFS= read -r -d '' file; do
    ((TOTAL++))
    update_artifact_file "$file"
    
    # Progress indicator
    if (( TOTAL % 10 == 0 )); then
        echo -n "."
    fi
done < <(find artifacts/definitions -name "*.yaml" -type f -print0)

echo
echo
echo "=== Update Complete ==="
echo "Total files processed: $TOTAL"
echo "Files updated: $UPDATED"
echo "Files skipped: $SKIPPED"
echo "Backup created: artifacts_backup_*.tar.gz"
echo
echo "Review changes with: git diff"
echo "Commit with: git add -A && git commit -m 'Bulk update artifact definitions with CyFir branding'"