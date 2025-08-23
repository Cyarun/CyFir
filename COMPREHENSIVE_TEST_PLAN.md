# Comprehensive Test Plan - CyFir Rebranding

## Test Overview
This plan covers all changes made in Phases 2 and 3 of the CyFir rebranding effort.

## Summary of Changes to Test
- **Environment variable compatibility** (VELOCIRAPTOR_* and CYFIR_*)
- **Binary name compatibility** (velociraptor and cyfir)
- **17 string updates** in various locations
- **Branding architecture** (prepared but not activated)

## Test Categories

### 1. Basic Functionality Tests ✅

#### 1.1 Binary Execution
```bash
# Test both binaries work
./velociraptor version
./cyfir version

# Expected: Both show identical output with name: cyfir
```

#### 1.2 Environment Variables
```bash
# Test old variables
VELOCIRAPTOR_CONFIG=test_config.yaml ./cyfir version
VELOCIRAPTOR_LITERAL_CONFIG='{"Client":{}}' ./cyfir version
VELOCIRAPTOR_API_CONFIG=api_config.yaml ./cyfir version

# Test new variables
CYFIR_CONFIG=test_config.yaml ./cyfir version
CYFIR_LITERAL_CONFIG='{"Client":{}}' ./cyfir version
CYFIR_API_CONFIG=api_config.yaml ./cyfir version

# Test precedence (CYFIR should win)
VELOCIRAPTOR_CONFIG=old.yaml CYFIR_CONFIG=new.yaml ./cyfir config show
```

### 2. Configuration Generation Tests

#### 2.1 Interactive Config Generation
```bash
./cyfir config generate --interactive

# Expected outputs:
# - "Welcome to the CyFir configuration generator"
# - "CyFir will store all files"
# - "CyFir will write logs to this directory"
```

#### 2.2 Basic Config Generation
```bash
./cyfir config generate > test_server.yaml
./cyfir config client --config test_server.yaml > test_client.yaml

# Verify configs are valid
./cyfir --config test_server.yaml config validate
```

### 3. String Update Verification

#### 3.1 Help Text
```bash
./cyfir --help | grep -i cyfir
./cyfir config --help | grep -i cyfir
```

#### 3.2 Artifact Descriptions
```bash
# Check artifact help shows updated descriptions
./cyfir artifacts show Generic.Detection.HashHunter | grep "CyFir accessor"
./cyfir artifacts show Windows.Detection.BinaryHunter | grep "CyFir accessor"
./cyfir artifacts show Server.Utils.CreateCollector | grep "CyFir banner"
```

### 4. Server/Client Tests

#### 4.1 Server Startup
```bash
# Generate test config
./cyfir config generate --self_signed > server_test.yaml

# Start server (in background or separate terminal)
./cyfir --config server_test.yaml frontend -v &
SERVER_PID=$!

# Wait for startup
sleep 5

# Check server is running
curl -k https://localhost:8000/server.pem

# Stop server
kill $SERVER_PID
```

#### 4.2 Client Connection
```bash
# Generate client config
./cyfir config client --config server_test.yaml > client_test.yaml

# Test client can show config
./cyfir --config client_test.yaml config show
```

### 5. Legacy Detection Tests

#### 5.1 Detection Function Tests (if on appropriate platform)
```bash
# Create test script to verify detection logic
cat > test_legacy_detection.go << 'EOF'
package main

import (
    "fmt"
    "github.com/Cyarun/CyFir/config"
)

func main() {
    detected := config.DetectLegacyInstallation()
    fmt.Printf("Legacy installation detected: %v\n", detected)
}
EOF

# This would need proper compilation, just documenting the test
```

### 6. Regression Tests

#### 6.1 Existing Features Still Work
```bash
# VQL queries work
./cyfir query "SELECT * FROM info()"

# Artifact listing works
./cyfir artifacts list | head -10

# Help system works
./cyfir help
```

#### 6.2 No Breaking Changes
```bash
# Old binary name works
./velociraptor version

# Old env vars work
VELOCIRAPTOR_CONFIG=test_config.yaml ./velociraptor version
```

### 7. Performance Tests

#### 7.1 Startup Time
```bash
# Measure startup time
time ./cyfir version
time ./velociraptor version

# Should be identical
```

#### 7.2 Memory Usage
```bash
# Monitor memory during basic operations
/usr/bin/time -v ./cyfir query "SELECT * FROM info()"
```

### 8. Documentation Tests

#### 8.1 README Files
```bash
# Verify documentation updates
grep -i cyfir tools/survey/README.md
grep -i cyfir branding/README.md
```

### 9. Build System Tests

#### 9.1 Compilation
```bash
# Clean build
go clean -cache
go build -tags "server_vql" -o test_build ./bin/

# Verify binary works
./test_build version
rm test_build
```

## Test Execution Checklist

### Phase 1: Quick Smoke Tests (5 minutes)
- [ ] Both binaries execute
- [ ] Version command works
- [ ] Basic help works
- [ ] No compilation errors

### Phase 2: Functionality Tests (15 minutes)
- [ ] Environment variables work (both old and new)
- [ ] Config generation works
- [ ] String updates visible in output
- [ ] No error messages mentioning Velociraptor where CyFir expected

### Phase 3: Integration Tests (30 minutes)
- [ ] Server starts successfully
- [ ] Client can connect
- [ ] Artifacts can be listed
- [ ] VQL queries work

### Phase 4: Compatibility Tests (15 minutes)
- [ ] Old binary name works
- [ ] Old environment variables work
- [ ] Mixed usage works (old binary, new env var)

## Success Criteria

### Must Pass:
1. All binaries execute without errors
2. Both old and new environment variables work
3. No functionality is broken
4. Updated strings appear in user-facing output

### Should Pass:
1. Performance is unchanged
2. All regression tests pass
3. Documentation is consistent

### Nice to Have:
1. Legacy detection functions work (platform-specific)
2. Build system works smoothly

## Known Issues/Limitations

1. **Protobuf generation blocked** - Branding configuration not fully integrated
2. **Some strings unchanged** - Registry paths, firewall rules kept for compatibility
3. **Build artifacts missing** - b0x.yaml files not in repository

## Test Results Template

```markdown
## Test Run: [DATE]

### Environment:
- OS: 
- Go Version: 
- Git Commit: 

### Results:
- Quick Smoke Tests: PASS/FAIL
- Functionality Tests: PASS/FAIL
- Integration Tests: PASS/FAIL
- Compatibility Tests: PASS/FAIL

### Issues Found:
1. 
2. 

### Overall Status: PASS/FAIL
```

## Next Steps After Testing

If all tests pass:
1. Deploy to staging environment
2. Run extended 24-hour test
3. Plan next phase of updates

If tests fail:
1. Document failures
2. Fix issues
3. Re-run failed tests
4. Update test plan if needed