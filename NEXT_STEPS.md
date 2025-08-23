# Recommended Next Steps for CyFir Rebranding

## Step 1: Test Current Changes (TODAY)
**Risk: LOW | Impact: Diagnostic only**

```bash
# 1.1 Clean build to ensure fresh state
make clean

# 1.2 Build with current changes
make auto

# 1.3 Run basic tests
make test_light

# 1.4 Test environment variable compatibility
chmod +x test_env_compat.sh
./test_env_compat.sh
```

**Expected Outcome**: All tests pass, confirming env var compatibility works

## Step 2: Implement Binary Name Compatibility (AFTER TESTS PASS)
**Risk: LOW | Impact: User-facing but backward compatible**

```bash
# 2.1 Create compatibility script
cat > create_cyfir_binary.sh << 'EOF'
#!/bin/bash
if [ -f "output/velociraptor" ]; then
    cp output/velociraptor output/cyfir
    echo "Created cyfir binary"
fi
EOF

chmod +x create_cyfir_binary.sh

# 2.2 Update Makefile to automatically create both binaries
# Add to the build targets
```

**Testing Required**:
- Both binaries work identically
- All commands function with both names
- No hardcoded "velociraptor" binary references break

## Step 3: Update Safe String Constants (AFTER BINARY COMPAT)
**Risk: LOW-MEDIUM | Impact: Log messages and display strings**

Target areas:
1. Log messages that don't affect parsing
2. User-facing messages in GUI
3. Help text and descriptions
4. Comments and documentation

**Testing Required**:
- Full regression test suite
- GUI functionality test
- Log parsing tools still work

## Step 4: Create Version Transition Strategy
**Risk: MEDIUM | Impact: Version reporting**

Consider adding a compatibility layer:
```go
// In version reporting
fmt.Sprintf("%s (formerly Velociraptor) %s", "CyFir", VERSION)
```

This helps users understand the transition.

## Step 5: Document Migration Path
**Risk: LOW | Impact: User documentation**

Create:
- MIGRATION.md - How to migrate from Velociraptor to CyFir
- FAQ.md - Common questions about the rebranding
- COMPATIBILITY.md - What remains compatible

## STOP POINTS - Do NOT Proceed Without Discussion:

### 🛑 Before changing any of these:
1. **Service names** - Would break existing installations
2. **Certificate CN/SAN** - Would break TLS
3. **Protocol magic bytes** - Would break client/server comms
4. **Database schemas** - Would require migration
5. **API endpoints** - Would break integrations
6. **Config file formats** - Would break existing configs
7. **Artifact namespaces** - Would break artifact compatibility

## Immediate Action Items:

### A. Run Tests Now
```bash
# Start with this - it's completely safe
make clean && make auto
```

### B. Create Rollback Point
```bash
# Tag current state for easy rollback
git tag -a pre-phase2-testing -m "Before Phase 2 testing"
```

### C. Monitor First Test Run
Run the test and capture output:
```bash
./test_env_compat.sh 2>&1 | tee test_results_$(date +%Y%m%d_%H%M%S).log
```

## Success Criteria Before Next Phase:

✅ All existing tests pass
✅ Environment variable compatibility confirmed
✅ No performance regression
✅ Binary runs without errors
✅ Can load existing config files

## Risk Mitigation:

1. **Test in isolated environment first**
2. **Keep detailed logs of all changes**
3. **Have rollback plan ready**
4. **Don't deploy to production until fully tested**
5. **Consider feature flags for gradual rollout**

---

**REMEMBER**: The goal is zero breakage. If anything fails, STOP and reassess.