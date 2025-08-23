#!/bin/bash
# Test script for environment variable compatibility

echo "===== CyFir Environment Variable Compatibility Test ====="
echo

# Create a minimal test config
cat > test_config.yaml << EOF
version:
  name: CyFir
  version: 0.75.1
Client:
  server_urls:
    - https://localhost:8000/
EOF

echo "1. Testing build..."
make auto
if [ $? -ne 0 ]; then
    echo "ERROR: Build failed!"
    exit 1
fi
echo "✓ Build successful"
echo

echo "2. Testing with no environment variables..."
./output/velociraptor version 2>/dev/null
if [ $? -eq 0 ]; then
    echo "✓ Binary runs without config"
else
    echo "✗ Binary failed without config"
fi
echo

echo "3. Testing with old VELOCIRAPTOR_CONFIG..."
export VELOCIRAPTOR_CONFIG=test_config.yaml
./output/velociraptor config show 2>&1 | grep -q "CyFir"
if [ $? -eq 0 ]; then
    echo "✓ VELOCIRAPTOR_CONFIG works"
else
    echo "✗ VELOCIRAPTOR_CONFIG failed"
fi
unset VELOCIRAPTOR_CONFIG
echo

echo "4. Testing with new CYFIR_CONFIG..."
export CYFIR_CONFIG=test_config.yaml
./output/velociraptor config show 2>&1 | grep -q "CyFir"
if [ $? -eq 0 ]; then
    echo "✓ CYFIR_CONFIG works"
else
    echo "✗ CYFIR_CONFIG failed"
fi
unset CYFIR_CONFIG
echo

echo "5. Testing with both set (CYFIR should take precedence)..."
# Create a second config to test precedence
cat > test_config2.yaml << EOF
version:
  name: CyFir-Priority
  version: 0.75.1
Client:
  server_urls:
    - https://localhost:8001/
EOF

export VELOCIRAPTOR_CONFIG=test_config.yaml
export CYFIR_CONFIG=test_config2.yaml
./output/velociraptor config show 2>&1 | grep -q "CyFir-Priority"
if [ $? -eq 0 ]; then
    echo "✓ CYFIR_CONFIG takes precedence correctly"
else
    echo "✗ Precedence test failed"
fi
unset VELOCIRAPTOR_CONFIG
unset CYFIR_CONFIG
echo

echo "6. Testing literal config..."
export CYFIR_LITERAL_CONFIG='{"version":{"name":"CyFir-Literal"}}'
./output/velociraptor config show 2>&1 | grep -q "CyFir-Literal"
if [ $? -eq 0 ]; then
    echo "✓ CYFIR_LITERAL_CONFIG works"
else
    echo "✗ CYFIR_LITERAL_CONFIG failed"
fi
unset CYFIR_LITERAL_CONFIG

# Cleanup
rm -f test_config.yaml test_config2.yaml

echo
echo "===== Test Complete ====="