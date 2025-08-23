# Phase 3: Branding Configuration Integration Plan

## Overview
The branding configuration system exists but is not yet integrated. This provides an excellent opportunity to implement a clean, configurable approach to the remaining rebranding work.

## Current Situation
- ✅ `config/branding.go` exists with full branding support
- ❌ Not integrated into the application
- ❌ Service names are hardcoded
- ❌ No legacy mode detection

## Integration Strategy

### Step 1: Add Branding to Main Config
First, we need to add branding settings to the main configuration:

```go
// In config/proto/config.proto or config.go
type Config struct {
    // ... existing fields ...
    
    // Branding configuration
    UseLegacyBranding bool `json:"use_legacy_branding"`
}
```

### Step 2: Detect Legacy Installations
Create detection logic:

```go
func DetectLegacyInstallation() bool {
    // Check for existing service names
    // Check for existing registry keys
    // Check for existing file paths
    // Return true if Velociraptor installation detected
}
```

### Step 3: Integration Points

#### 3.1 Service Installation
- `bin/installer_windows.go`
- `bin/service_compat.go`
- Update to use `GetBranding()` for service names

#### 3.2 Configuration Generation
- `tools/survey/server.go`
- `config/config.go`
- Use branding for default values

#### 3.3 Display Strings
- Replace hardcoded "CyFir" with `branding.DisplayName`
- Replace hardcoded company names with `branding.Company`

### Step 4: Command Line Flag
Add a flag to force legacy mode:

```bash
./cyfir --legacy-branding config generate
```

## Implementation Order

### Phase 3A: Wire Up Configuration
1. Add branding field to config
2. Update config loading to set branding mode
3. Add command line flag support
4. Test configuration loading

### Phase 3B: Service Name Integration
1. Update service installation code
2. Update service compatibility functions
3. Test on Windows/Linux/Mac
4. Verify upgrade scenarios

### Phase 3C: Display String Integration
1. Find all hardcoded display strings
2. Replace with branding configuration calls
3. Update help text and messages
4. Test UI appearance

## Benefits of This Approach

1. **Clean Migration**: Users can toggle between brands
2. **Backward Compatible**: Existing installations work
3. **Centralized Control**: One place to manage branding
4. **Safe Rollout**: Can test both modes easily
5. **Future Proof**: Easy to update branding later

## Risk Mitigation

### Testing Matrix
- [ ] Fresh install with CyFir branding
- [ ] Fresh install with legacy branding
- [ ] Upgrade from Velociraptor to CyFir
- [ ] Toggle branding mode on existing install
- [ ] Service management with both modes

### Rollback Plan
- Configuration flag allows instant rollback
- No hardcoded changes to reverse
- Git tags for code rollback if needed

## Example Implementation

### Update Service Installation
```go
// Before:
name := "CyFir"

// After:
branding := GetBranding(config_obj.UseLegacyBranding)
name := branding.WindowsServiceName
```

### Update Display Strings
```go
// Before:
fmt.Println("Welcome to CyFir")

// After:
branding := GetBranding(config_obj.UseLegacyBranding)
fmt.Printf("Welcome to %s\n", branding.DisplayName)
```

## Success Criteria
1. Single configuration controls all branding
2. Can switch between brands via config
3. Existing installations continue working
4. New installations use new branding
5. All 318 remaining references addressed

## Time Estimate
- Phase 3A: 2-4 hours (configuration wiring)
- Phase 3B: 4-6 hours (service integration)
- Phase 3C: 6-8 hours (display strings)
- Testing: 4-6 hours

Total: ~20-30 hours for complete integration

This approach is much cleaner and safer than manual string replacement!