# CyFir Rebranding Test Plan

## Overview
This document outlines the comprehensive test plan to validate that all rebranding changes work correctly and maintain backward compatibility.

## Test Categories

### 1. Unit Tests
- **Command**: `make test_light`
- **Purpose**: Run all unit tests without race detection
- **Expected**: All tests should pass

### 2. Race Condition Tests
- **Command**: `make test`
- **Purpose**: Run tests with race detection enabled
- **Expected**: No race conditions detected

### 3. Configuration Tests
- **Manual Test**: Test both old and new environment variables
  - `VELOCIRAPTOR_CONFIG` should still work
  - `CYFIR_CONFIG` should take precedence
  - Binary name compatibility (velociraptor → cyfir)

### 4. Build Tests
- **Command**: `make` or `make linux`
- **Purpose**: Ensure the binary builds successfully
- **Expected**: Binary created as `output/cyfir`

### 5. GUI Tests
- **Manual Test**: Start the GUI and verify:
  - Login page shows "CyFir Login"
  - Theme names show "CyFir (light/dark)"
  - All user-visible text shows CyFir
  - Tool upload still works with VelociraptorWindows tools

### 6. Artifact Tests
- **Command**: `make golden`
- **Purpose**: Run golden file tests for artifacts
- **Expected**: All artifact tests pass

### 7. Backward Compatibility Tests
- **Service Names**: Verify old service names still work
  - Windows: "Velociraptor"
  - Linux: "velociraptor_client"
- **Tool Names**: Verify artifacts can still use:
  - VelociraptorWindows
  - VelociraptorLinux
  - VelociraptorCollector

### 8. Integration Tests
- Start server with old config
- Connect client with old binary name
- Run basic collection
- Verify results appear correctly

## Test Execution Steps

1. **Clean Build**
   ```bash
   make clean
   make
   ```

2. **Run Unit Tests**
   ```bash
   make test_light
   ```

3. **Run Race Tests** (if time permits)
   ```bash
   make test
   ```

4. **Test Configuration Loading**
   ```bash
   # Test old env var
   export VELOCIRAPTOR_CONFIG=/path/to/config.yaml
   ./output/cyfir config show
   
   # Test new env var (should override)
   export CYFIR_CONFIG=/path/to/config.yaml
   ./output/cyfir config show
   ```

5. **Test Binary Compatibility**
   ```bash
   # Create symlink for backward compatibility
   ln -s ./output/cyfir ./output/velociraptor
   ./output/velociraptor version
   ```

## Expected Results

### ✅ Pass Criteria
- All unit tests pass
- Binary builds successfully
- Both old and new env vars work
- GUI shows CyFir branding
- Artifacts execute correctly
- Backward compatibility maintained

### ❌ Fail Criteria
- Any unit test failures
- Build failures
- Environment variables not recognized
- GUI showing Velociraptor text
- Artifacts failing to load tools
- Service startup failures

## Risk Areas

1. **Tool Loading**: Artifacts still reference VelociraptorWindows
2. **Service Names**: Windows/Linux service compatibility
3. **Import Paths**: Changed from velocidex.com to github.com/Cyarun/CyFir
4. **Configuration**: Protobuf field for branding not generated

## Notes
- Focus on user-visible functionality
- Document any failures for v2.0 planning
- Backward compatibility is critical