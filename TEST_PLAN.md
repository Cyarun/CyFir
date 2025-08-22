# CyFir Rebranding Test Plan

## Phase 1 Tests (Completed Changes)
- [x] Module imports work correctly
- [x] Build completes successfully
- [ ] Unit tests pass
- [ ] Binary starts without errors

## Phase 2 Tests (Compatibility Layer)

### Environment Variable Tests
1. Test with VELOCIRAPTOR_CONFIG set
2. Test with CYFIR_CONFIG set  
3. Test with both set (CYFIR should take precedence)
4. Test with neither set

### Binary Name Tests
1. velociraptor binary works as before
2. cyfir binary works identically
3. All command line flags work with both

### Service Tests
1. Windows service installation with old name
2. Systemd service with old name
3. Configuration files still load

### Client-Server Tests
1. Old client connects to new server
2. New client connects to old server
3. Mixed deployment works

## Test Commands

```bash
# Basic functionality test
./output/velociraptor config show
./output/cyfir config show

# Environment variable test
VELOCIRAPTOR_CONFIG=/path/to/config ./output/velociraptor config show
CYFIR_CONFIG=/path/to/config ./output/cyfir config show

# Service test (Linux)
sudo ./output/velociraptor service install
sudo systemctl status velociraptor

# Unit tests
go test -v --tags server_vql ./...