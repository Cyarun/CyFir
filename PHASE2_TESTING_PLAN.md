# Phase 2 Comprehensive Testing Plan

## Test Environment Setup
1. Create isolated test environment
2. Document system specs and versions
3. Prepare rollback procedures

## Phase 2A Tests (Environment Variables) ✅
Already tested basic functionality. Need to verify:

### 1. Environment Variable Tests
```bash
# Test 1: Old variable still works
VELOCIRAPTOR_CONFIG=test_complete_config.yaml ./cyfir version

# Test 2: New variable works
CYFIR_CONFIG=test_complete_config.yaml ./cyfir version

# Test 3: Precedence (CYFIR should win)
VELOCIRAPTOR_CONFIG=config1.yaml CYFIR_CONFIG=config2.yaml ./cyfir config show

# Test 4: API config
CYFIR_API_CONFIG=api_config.yaml ./cyfir api-client ...

# Test 5: Literal config
CYFIR_LITERAL_CONFIG='{"Client":{"server_urls":["https://localhost:8000/"]}}' ./cyfir version
```

### 2. Binary Compatibility Tests
```bash
# Both binaries should work identically
./velociraptor version
./cyfir version

# All subcommands should work
./velociraptor config generate > test1.yaml
./cyfir config generate > test2.yaml
diff test1.yaml test2.yaml  # Should be identical
```

## Phase 2B-1 Tests (String Changes)

### 1. Setup/Survey Tool Tests
```bash
# Test config generation shows new messages
./cyfir config generate --interactive

# Expected: Should see "Welcome to the CyFir configuration generator"
# NOT: "Welcome to the Velociraptor configuration generator"
```

### 2. Accessor Tests
```bash
# Test me accessor still works
./cyfir query "SELECT * FROM glob(accessor='me', globs='/**')" 

# Check description
./cyfir query "SELECT * FROM info()" | grep -i "bundled inside"
# Should show: "bundled inside the CyFir binary"
```

### 3. Comment Verification
```bash
# Comments don't affect runtime, but verify files still compile
go test -v --tags server_vql ./vql/server/downloads/...
go test -v --tags server_vql ./vql/tools/...
go test -v --tags server_vql ./accessors/zip/...
```

## Integration Tests

### 1. Server Startup Test
```bash
# Generate test server config
./cyfir config generate > server.config.yaml

# Start server (in background or separate terminal)
./cyfir --config server.config.yaml frontend -v

# Check server is running
curl -k https://localhost:8000/app/index.html
```

### 2. Client Connection Test
```bash
# Generate client config from server
./cyfir config client --config server.config.yaml > client.config.yaml

# Test client can connect
./cyfir --config client.config.yaml client -v
```

### 3. Basic Artifact Collection
```bash
# Run a simple artifact
./cyfir --config server.config.yaml query "SELECT * FROM info()"

# Run built-in artifact
./cyfir --config server.config.yaml artifacts collect Generic.Client.Info
```

## Performance Tests

### 1. Memory Usage
```bash
# Monitor memory before/after changes
/usr/bin/time -v ./cyfir --config server.config.yaml frontend &
# Let it run for 5 minutes, then check memory stats
```

### 2. Startup Time
```bash
# Measure startup time
time ./cyfir version
time ./cyfir --config server.config.yaml frontend --help
```

## Regression Tests

### 1. Existing Config Compatibility
- Load configs created with old version
- Ensure all features still work
- Check for any deprecation warnings

### 2. Data Format Tests
- Verify data written is same format
- Check database compatibility
- Ensure file uploads work

## GUI Tests (if GUI available)

### 1. Web Interface
- Login works
- Navigation functions
- No broken links or references
- Check for any "Velociraptor" text in UI

### 2. API Tests
```bash
# Test API endpoints
curl -k -X GET https://localhost:8000/api/v1/GetVersion
```

## Error Scenarios

### 1. Missing Config
```bash
./cyfir client  # Should show helpful error
```

### 2. Invalid Config
```bash
echo "invalid: yaml: content" > bad.yaml
./cyfir --config bad.yaml frontend  # Should show clear error
```

## Test Report Template

### Test Results Summary
- [ ] All Phase 2A tests pass
- [ ] All Phase 2B-1 tests pass
- [ ] No performance regression
- [ ] No memory leaks
- [ ] GUI functions correctly
- [ ] API compatibility maintained

### Issues Found
1. Issue: ___________
   - Severity: High/Medium/Low
   - Resolution: ___________

### Sign-off Checklist
- [ ] Code compiles without warnings
- [ ] All tests pass
- [ ] No regressions identified
- [ ] Performance acceptable
- [ ] Ready for Phase 2B-2

## Rollback Procedures

If any critical issues found:
```bash
# Rollback to last known good
git reset --hard phase2a-complete

# Or rollback to before Phase 2
git reset --hard master~10

# Rebuild
go build -tags "server_vql" -o cyfir ./bin/
```

## Next Steps After Testing

If all tests pass:
1. Document test results
2. Get sign-off
3. Proceed to Phase 2B-2 (more string updates)

If issues found:
1. Document failures
2. Create fixes
3. Re-test affected areas
4. Update test plan