#!/bin/bash
# Create compatibility symlinks for smooth migration

echo "Creating compatibility symlinks..."

# Check if cyfir binary exists
if [ -f "output/cyfir" ]; then
    # Create velociraptor symlink for compatibility
    ln -sf cyfir output/velociraptor
    echo "✓ Created output/velociraptor -> cyfir"
fi

# For Windows builds
if [ -f "output/cyfir.exe" ]; then
    ln -sf cyfir.exe output/velociraptor.exe
    echo "✓ Created output/velociraptor.exe -> cyfir.exe"
fi

# For production builds
for file in output/cyfir-*; do
    if [ -f "$file" ]; then
        base=$(basename "$file")
        velo_name=${base/cyfir/velociraptor}
        ln -sf "$base" "output/$velo_name"
        echo "✓ Created output/$velo_name -> $base"
    fi
done

echo "Compatibility links created successfully!"