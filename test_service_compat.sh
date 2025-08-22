#!/bin/bash

echo "=== Testing Service Name Compatibility ==="

# Test service installation with legacy name
echo "1. Testing legacy service name..."
./output/velociraptor service install --help | grep -i service

echo ""
echo "2. Testing new binary with service command..."
./output/cyfir service install --help | grep -i service

echo ""
echo "3. Checking config generation..."
./output/cyfir config generate > test_config.yaml 2>&1
if [ $? -eq 0 ]; then
    echo "✓ Config generation works"
    
    # Check service names in config
    echo "4. Service names in generated config:"
    grep -i "service" test_config.yaml | head -5
else
    echo "✗ Config generation failed"
fi

echo ""
echo "=== Service Compatibility Test Complete ==="