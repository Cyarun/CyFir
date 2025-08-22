# CyFir Rebranding Safety Checklist

## Pre-Build Verification ✓

### Dependencies Status:
- [x] Go 1.24.4 installed (exceeds minimum 1.23.2)
- [x] Essential build tools (make, gcc, git) installed
- [x] Go modules verified and up to date
- [x] Node.js and npm available for GUI
- [x] fileb0x available for asset embedding

### Code Changes Made (Safe):
1. **Module Path**: Changed from `www.velocidex.com/golang/velociraptor` to `github.com/Cyarun/CyFir`
2. **Import Paths**: Updated in 5491 files
3. **Copyright**: Updated to CynorSense Solutions Pvt. Ltd.
4. **Documentation URLs**: Updated to cyfir.cynorsense.com
5. **Constants**: Updated USER_AGENT to "CyFir"

### What We DID NOT Change (Intentionally):
1. **Binary name in build system**: Still "velociraptor" 
2. **Service names**: Unchanged to avoid breaking installations
3. **Protocol identifiers**: Unchanged for compatibility
4. **Environment variable names**: Still VELOCIRAPTOR_* (created wrapper for CYFIR_*)
5. **Artifact definitions**: Unchanged
6. **Client-server communication protocols**: Unchanged

### Compatibility Measures:
1. Created `cyfir` wrapper script that translates CYFIR_* env vars
2. Created `build_compat.sh` to build both binary names
3. Maintained all original functionality

## Safe Test Sequence:

### Step 1: Minimal Build Test
```bash
# Build without YARA or extras for quick test
make linux_bare
```

### Step 2: Version Check
```bash
# Should show version and build info
./output/velociraptor-*-linux-amd64 version
```

### Step 3: Config Test
```bash
# Test configuration loading
./output/velociraptor-*-linux-amd64 config show
```

### Step 4: Unit Tests (if build works)
```bash
# Run basic tests without race detection
make test_light
```

### Step 5: Full Build (if tests pass)
```bash
# Build with all features
make auto
```

## Rollback Plan:

If anything breaks:
1. Git status to see all changes
2. Key files to check:
   - go.mod (module path)
   - constants/constants.go (branding)
   - All import statements

## Next Safe Steps After Testing:

1. **Create tagged release**: Tag current state before more changes
2. **Document all changes**: Update migration guide
3. **Test deployment scenarios**: 
   - Fresh install
   - Upgrade from Velociraptor
   - Mixed environment

## DO NOT PROCEED WITH:
- Changing binary names in core build system
- Modifying service installation scripts
- Changing network protocols
- Renaming artifacts
- Modifying client-server handshake

Until extensive testing is complete!