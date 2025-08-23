#!/bin/bash
# Create cyfir binary alongside velociraptor for compatibility

echo "Creating binary compatibility..."

# For development binary
if [ -f "test_binary" ]; then
    cp test_binary velociraptor
    cp test_binary cyfir
    echo "✓ Created velociraptor and cyfir binaries from test_binary"
fi

# For output directory
if [ -f "output/velociraptor" ]; then
    cp output/velociraptor output/cyfir
    echo "✓ Created output/cyfir from output/velociraptor"
fi

# Test both work
echo
echo "Testing binaries:"
if [ -f "velociraptor" ]; then
    echo -n "velociraptor: "
    ./velociraptor version | grep "name:" || echo "FAILED"
fi

if [ -f "cyfir" ]; then
    echo -n "cyfir: "
    ./cyfir version | grep "name:" || echo "FAILED"
fi