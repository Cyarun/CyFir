#!/bin/bash
# Update remaining Go file references to CyFir

echo "=== Updating Go File References ==="

# Update string literals in Go files
echo "Updating string literals..."

# Update user-facing strings
find . -name "*.go" -type f -exec grep -l '"Velociraptor' {} \; | while read file; do
    echo "Checking: $file"
    sed -i 's/"Velociraptor server"/"CyFir server"/g' "$file"
    sed -i 's/"Velociraptor client"/"CyFir client"/g' "$file"
    sed -i 's/"Velociraptor agent"/"CyFir agent"/g' "$file"
    sed -i 's/"Velociraptor GUI"/"CyFir GUI"/g' "$file"
    sed -i 's/"Velociraptor API"/"CyFir API"/g' "$file"
    sed -i 's/"Velociraptor service"/"CyFir service"/g' "$file"
    sed -i 's/"Velociraptor instance"/"CyFir instance"/g' "$file"
done

# Update error messages
echo "Updating error messages..."
find . -name "*.go" -type f -exec grep -l 'Error.*Velociraptor' {} \; | while read file; do
    echo "Updating errors in: $file"
    sed -i 's/Velociraptor server/CyFir server/g' "$file"
    sed -i 's/Velociraptor client/CyFir client/g' "$file"
done

# Update log messages
echo "Updating log messages..."
find . -name "*.go" -type f -exec grep -l 'log.*Velociraptor' {} \; | while read file; do
    echo "Updating logs in: $file"
    sed -i 's/Starting Velociraptor/Starting CyFir/g' "$file"
    sed -i 's/Stopping Velociraptor/Stopping CyFir/g' "$file"
    sed -i 's/Velociraptor started/CyFir started/g' "$file"
    sed -i 's/Velociraptor stopped/CyFir stopped/g' "$file"
done

# Update comments (careful not to break code)
echo "Updating comments..."
find . -name "*.go" -type f -exec grep -l '// .*Velociraptor' {} \; | while read file; do
    echo "Updating comments in: $file"
    sed -i 's|// Velociraptor|// CyFir|g' "$file"
    sed -i 's|// The Velociraptor|// The CyFir|g' "$file"
    sed -i 's|// This Velociraptor|// This CyFir|g' "$file"
done

echo
echo "=== Update Complete ==="
echo "Review changes carefully with: git diff"
echo "Some references may need manual review to ensure functionality isn't broken."