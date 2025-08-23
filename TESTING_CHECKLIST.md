# CyFir Rebranding Testing Checklist

## Pre-Testing Setup
- [ ] Create backup of current working state
- [ ] Document current git commit hash
- [ ] Ensure clean working directory

## Phase 1: Build Tests
- [ ] `make clean` completes successfully
- [ ] `make auto` builds without errors
- [ ] `make test_light` passes all tests
- [ ] Binary size is similar to before changes
- [ ] No new compiler warnings

## Phase 2: Environment Variable Tests
- [ ] Run test_env_compat.sh script
- [ ] Old VELOCIRAPTOR_CONFIG still works
- [ ] New CYFIR_CONFIG works
- [ ] CYFIR_CONFIG takes precedence when both set
- [ ] VELOCIRAPTOR_LITERAL_CONFIG still works
- [ ] CYFIR_LITERAL_CONFIG works
- [ ] VELOCIRAPTOR_API_CONFIG still works
- [ ] CYFIR_API_CONFIG works

## Phase 3: Binary Compatibility Tests
- [ ] `velociraptor version` shows correct version
- [ ] `cyfir version` shows same version
- [ ] All subcommands work with both binaries
- [ ] Help text displays correctly

## Phase 4: Configuration Tests
- [ ] Existing config files load without errors
- [ ] Server starts with existing config
- [ ] Client connects with existing config
- [ ] Web UI loads correctly

## Phase 5: Service Tests (if applicable)

### Linux Service Tests
- [ ] Install as systemd service
- [ ] Service starts successfully
- [ ] Service stops cleanly
- [ ] Logs show no errors
- [ ] Service survives reboot

### Windows Service Tests
- [ ] Install as Windows service
- [ ] Service starts successfully
- [ ] Service stops cleanly
- [ ] Event logs show no errors
- [ ] Service survives reboot

## Phase 6: Client-Server Tests
- [ ] Server accepts connections
- [ ] Client can enroll
- [ ] Artifacts can be collected
- [ ] File uploads work
- [ ] Hunt creation works
- [ ] Results are stored correctly

## Phase 7: Data Compatibility Tests
- [ ] Old data can be read
- [ ] New data format is compatible
- [ ] Database queries work
- [ ] File store access works
- [ ] Indexes are maintained

## Phase 8: Integration Tests
- [ ] API endpoints respond correctly
- [ ] gRPC connections work
- [ ] Third-party integrations work
- [ ] Monitoring/alerting works

## Phase 9: Performance Tests
- [ ] Memory usage is similar
- [ ] CPU usage is similar
- [ ] Startup time is similar
- [ ] Query performance unchanged

## Phase 10: Rollback Test
- [ ] Document rollback procedure
- [ ] Test rollback to previous version
- [ ] Verify data compatibility after rollback
- [ ] Ensure no data loss

## Sign-off
- [ ] All tests passed
- [ ] No regressions identified
- [ ] Performance acceptable
- [ ] Documentation updated
- [ ] Ready for next phase

## Notes Section
Use this section to document any issues found, workarounds needed, or special considerations:

---

### Issue Log:
1. 
2. 
3. 

### Workarounds:
1. 
2. 
3.