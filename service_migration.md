# Service Name Migration Strategy

## Current Service Names:
- Windows: "Velociraptor" service
- Linux: "velociraptor" systemd service  
- macOS: "com.velocidex.velociraptor" launchd service

## Proposed Approach:

### 1. Dual Service Support
- Keep existing service names working
- Add new service names (CyFir) as aliases
- Both point to same binary

### 2. Installation Script Updates
```bash
# Linux systemd example
# Create both service files
/etc/systemd/system/velociraptor.service  # Legacy
/etc/systemd/system/cyfir.service         # New (symlink to legacy)

# Windows
# Register with both names or use display name change
```

### 3. Migration Path
1. New installs: Use "cyfir" service name
2. Upgrades: Keep existing name, add alias
3. Future: Deprecate old name after 6-12 months

### 4. Configuration Compatibility
- Service configs continue to work with either name
- No breaking changes to existing deployments