# CyFir Rebranding - Phase 2 Detailed Plan

## Current Status Assessment

### What We've Changed:
1. Git remote configuration
2. Go module path and imports
3. Some branding constants
4. Documentation URLs

### What Still References Velociraptor:
- 284 Go files with Velociraptor references
- Test files and test data
- Binary names in build system
- Many internal references

## CRITICAL SAFETY PRINCIPLES

1. **DO NOT change anything that affects protocol compatibility**
2. **DO NOT change data formats or storage**
3. **DO NOT change API contracts**
4. **ALWAYS maintain backward compatibility**
5. **TEST after EVERY change**

## Phase 2A: Safe Internal Changes (Low Risk)

### Step 1: Update Non-Critical String Constants
**Risk Level: LOW**
- Log messages that are informational only
- Comments in code
- Internal variable names that don't affect external interfaces

**Testing Required:**
```bash
make test_light
./output/velociraptor version
```

### Step 2: Create Compatibility Wrappers
**Risk Level: LOW**
- Add cyfir as an alias to velociraptor binary
- Support both VELOCIRAPTOR_* and CYFIR_* environment variables
- Keep all old names working

**Implementation:**
1. Update config loaders to check both env vars
2. Create symlinks or copies for binaries
3. Add compatibility functions for name lookups

**Testing Required:**
```bash
# Test both environment variables work
VELOCIRAPTOR_CONFIG=test.yaml ./output/velociraptor config show
CYFIR_CONFIG=test.yaml ./output/velociraptor config show

# Test both binaries work
./output/velociraptor version
./output/cyfir version
```

## Phase 2B: Medium Risk Changes (Requires Extensive Testing)

### Step 3: Update User-Visible Strings
**Risk Level: MEDIUM**
- Error messages shown to users
- GUI text and labels
- Help text and command descriptions

**Safety Measures:**
- Create feature flag to toggle between old/new strings
- Test with both flags enabled/disabled
- Get user feedback before making permanent

### Step 4: Update Service Names (Windows/Linux)
**Risk Level: MEDIUM-HIGH**
- Service display names only (not service IDs)
- Keep service IDs unchanged for compatibility
- Update descriptions

**Testing Required:**
- Full service lifecycle testing
- Upgrade testing from old to new version
- Multi-version deployment testing

## Phase 2C: High Risk Changes (Defer or Avoid)

### Items to DEFER:
1. **Protocol identifiers** - Could break client-server communication
2. **File format headers** - Would make old data unreadable
3. **API endpoint names** - Would break integrations
4. **Database schema** - Would require migrations
5. **Certificate CN fields** - Would break TLS verification

## Detailed Next Steps

### 1. Environment Variable Compatibility (TODAY)

Let's implement the compatibility layer properly:

```go
// config/env_loader.go
func LoadConfigFromEnvironment() (*Config, error) {
    // Check new variable first
    if path := os.Getenv("CYFIR_CONFIG"); path != "" {
        return LoadConfig(path)
    }
    // Fall back to old variable
    if path := os.Getenv("VELOCIRAPTOR_CONFIG"); path != "" {
        return LoadConfig(path)
    }
    return nil, ErrNoConfig
}
```

### 2. Binary Name Compatibility (TODAY)

Update the build system to create both binaries:

```makefile
build:
    go build -o output/velociraptor ./bin/
    cp output/velociraptor output/cyfir
```

### 3. Create Comprehensive Test Suite

```bash
#!/bin/bash
# test_compatibility.sh

echo "Testing environment variable compatibility..."
# Test cases for env vars

echo "Testing binary compatibility..."
# Test cases for binaries

echo "Testing service compatibility..."
# Test cases for services

echo "Testing client-server compatibility..."
# Test cases for protocol
```

## Testing Strategy

### Level 1: Unit Tests
```bash
go test -v --tags server_vql ./...
```

### Level 2: Integration Tests
- Start server with old config
- Connect client with new binary
- Run artifacts
- Check data collection

### Level 3: Deployment Tests
- Install as service
- Run for 24 hours
- Check logs for errors
- Monitor performance

## Decision Points

Before proceeding to each next phase, we need to verify:

1. All tests pass
2. No functionality is broken
3. Performance is unchanged
4. Existing deployments work
5. Data formats are compatible

## Roll-back Plan

If anything breaks:
1. Git reset to last known good commit
2. Rebuild with old code
3. Document what failed
4. Adjust plan accordingly

## Questions to Answer First:

1. Are there any production deployments we need to consider?
2. What's the timeline for this transition?
3. Are there any external integrations that might break?
4. Do we need to maintain compatibility with specific versions?