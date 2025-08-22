#!/bin/bash
# Safe string updates - only user-visible messages

echo "=== Safe String Updates for CyFir ==="
echo "This script only updates user-visible strings, not technical identifiers"
echo

# Counter
UPDATED=0
TOTAL=0

# Update log messages (safe - user visible)
echo "Updating log messages..."
find . -name "*.go" -type f -exec grep -l 'log.*".*Velociraptor' {} \; | while read file; do
    echo "Updating: $file"
    # Only update within quoted strings in log statements
    sed -i 's/\(log[^"]*"\)\([^"]*\)Velociraptor server/\1\2CyFir server/g' "$file"
    sed -i 's/\(log[^"]*"\)\([^"]*\)Velociraptor client/\1\2CyFir client/g' "$file"
    sed -i 's/\(log[^"]*"\)\([^"]*\)Starting Velociraptor/\1\2Starting CyFir/g' "$file"
    sed -i 's/\(log[^"]*"\)\([^"]*\)Stopping Velociraptor/\1\2Stopping CyFir/g' "$file"
    ((UPDATED++))
done

# Update fmt.Sprint* messages (safe - user visible)
echo
echo "Updating formatted output messages..."
find . -name "*.go" -type f -exec grep -l 'fmt\.S.*".*Velociraptor' {} \; | while read file; do
    echo "Updating: $file"
    sed -i 's/\(fmt\.S[^"]*"\)\([^"]*\)Velociraptor server/\1\2CyFir server/g' "$file"
    sed -i 's/\(fmt\.S[^"]*"\)\([^"]*\)Velociraptor client/\1\2CyFir client/g' "$file"
    ((UPDATED++))
done

# Update panic messages (safe - user visible)
echo
echo "Updating panic messages..."
find . -name "*.go" -type f -exec grep -l 'panic.*".*Velociraptor' {} \; | while read file; do
    echo "Updating: $file"
    sed -i 's/\(panic[^"]*"\)\([^"]*\)Velociraptor/\1\2CyFir/g' "$file"
    ((UPDATED++))
done

# Update comments that describe user-facing features
echo
echo "Updating user-facing comments..."
find . -name "*.go" -type f -exec grep -l '// .*Velociraptor server\|// .*Velociraptor client' {} \; | while read file; do
    echo "Updating comments in: $file"
    sed -i 's|// \(.*\)Velociraptor server|// \1CyFir server|g' "$file"
    sed -i 's|// \(.*\)Velociraptor client|// \1CyFir client|g' "$file"
    ((UPDATED++))
done

echo
echo "=== Safe Updates Complete ==="
echo "Files updated: $UPDATED"
echo
echo "NOT changed (intentionally):"
echo "- Function names"
echo "- Variable names"
echo "- API endpoints"
echo "- VQL references"
echo "- Package imports"
echo "- Protocol definitions"