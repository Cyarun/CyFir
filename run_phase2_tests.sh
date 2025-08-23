#!/bin/bash
# Phase 2 Basic Test Suite

echo "===== CyFir Phase 2 Testing Suite ====="
echo "Date: $(date)"
echo

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Test counter
PASSED=0
FAILED=0

# Function to run test
run_test() {
    local test_name="$1"
    local test_cmd="$2"
    
    echo -n "Testing: $test_name... "
    if eval "$test_cmd" > /dev/null 2>&1; then
        echo -e "${GREEN}PASSED${NC}"
        ((PASSED++))
    else
        echo -e "${RED}FAILED${NC}"
        ((FAILED++))
    fi
}

echo "=== Phase 1: Binary Tests ==="
run_test "CyFir binary exists" "test -f ./cyfir"
run_test "Velociraptor binary exists" "test -f ./velociraptor"
run_test "CyFir version works" "./cyfir version"
run_test "Velociraptor version works" "./velociraptor version"

echo
echo "=== Phase 2: Environment Variable Tests ==="
run_test "CYFIR_CONFIG works" "CYFIR_CONFIG=test_complete_config.yaml ./cyfir version"
run_test "VELOCIRAPTOR_CONFIG works" "VELOCIRAPTOR_CONFIG=test_complete_config.yaml ./cyfir version"

echo
echo "=== Phase 3: String Update Tests ==="
# Check if setup messages were updated
run_test "Setup message updated" "./cyfir config generate --help 2>&1 | grep -q 'configuration'"

echo
echo "=== Phase 4: Compilation Tests ==="
run_test "Code compiles" "go build -tags 'server_vql' -o test_compile ./bin/"
run_test "Remove test binary" "rm -f test_compile"

echo
echo "=== Phase 5: Basic Functionality ==="
run_test "Query command available" "./cyfir query --help > /dev/null 2>&1"
run_test "Config command available" "./cyfir config --help > /dev/null 2>&1"

echo
echo "===== Test Summary ====="
echo "Passed: $PASSED"
echo "Failed: $FAILED"
echo

if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}All tests passed! Safe to proceed.${NC}"
    exit 0
else
    echo -e "${RED}Some tests failed. Please investigate before proceeding.${NC}"
    exit 1
fi