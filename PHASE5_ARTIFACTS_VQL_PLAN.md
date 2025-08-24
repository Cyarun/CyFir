# Phase 5: Artifacts and VQL Updates

## Critical Discovery

While the core application shows CyFir branding, the artifacts (which users interact with heavily) still contain 81 Velociraptor references. This is a significant gap in the rebranding.

## High-Priority Artifact Updates Needed

### 1. Tool Names (CRITICAL)
The artifact `Server.Internal.ToolDependencies.yaml` defines tool names used throughout:
- `VelociraptorWindows`
- `VelociraptorLinux`
- `VelociraptorCollector`
- `VelociraptorWindowsMSI`

These names are referenced in many artifacts and changing them would break artifact compatibility.

### 2. Artifact Descriptions (HIGH)
Many artifacts have descriptions mentioning Velociraptor:
- `Admin.Client.Uninstall`: "Uninstall Velociraptor from the endpoint"
- `Windows.Forensics.BulkExtractor`: References Velociraptor capabilities
- Multiple upgrade artifacts reference Velociraptor

### 3. Parameter Defaults (MEDIUM)
Several artifacts have parameters defaulting to "Velociraptor":
- `Admin.Client.Uninstall`: DisplayNameRegex defaults to "Velociraptor"
- Various quarantine artifacts reference VelociraptorFrontEnd

## Compatibility Concerns

### Tool Name Problem
Changing tool names like `VelociraptorWindows` to `CyFirWindows` would:
- Break existing artifacts that reference these tools
- Require updating all artifacts simultaneously
- Break backward compatibility

### Possible Solutions:
1. **Alias System**: Support both old and new tool names
2. **Versioned Artifacts**: New artifact versions with CyFir names
3. **Keep Internal Names**: Only update user-visible descriptions

## Recommended Approach

### Phase 5A: Update Descriptions Only (Safe)
1. Update artifact descriptions to mention CyFir
2. Keep tool names unchanged for compatibility
3. Update help text and documentation

### Phase 5B: Create Transition Artifacts (Medium Risk)
1. Create new versions of key artifacts with CyFir branding
2. Keep old artifacts for compatibility
3. Mark old ones as deprecated

### Phase 5C: Tool Name Migration (High Risk - Future)
1. Implement tool name aliasing system
2. Support both VelociraptorWindows and CyFirWindows
3. Gradually transition over major version

## Immediate Actions Needed

### 1. Update Critical User-Facing Descriptions
```yaml
# Admin.Client.Uninstall
description: |
  Uninstall CyFir from the endpoint.
  
  This artifact uninstalls a CyFir client (or any other MSI
  package) from the endpoint.
```

### 2. Update Parameter Descriptions
Keep defaults for compatibility but update descriptions:
```yaml
- name: DisplayNameRegex
  type: regex
  default: Velociraptor  # Keep for compatibility
  description: A regex that will match the package to uninstall (default matches legacy Velociraptor installations).
```

### 3. Add Migration Notes
Add notes to artifacts explaining the transition:
```yaml
description: |
  This artifact handles both Velociraptor and CyFir installations
  for backward compatibility during the transition period.
```

## Impact Assessment

### User Experience Impact: HIGH
- Users run artifacts daily
- Artifact descriptions are highly visible
- Tool names appear in many places

### Risk Assessment: MEDIUM
- Changing descriptions is safe
- Changing tool names is risky
- Need careful compatibility planning

## Recommendation

The rebranding is NOT complete without updating artifacts. This requires:
1. Immediate update of artifact descriptions (1-2 hours)
2. Careful planning for tool name transition
3. Possible major version release for breaking changes

This is a significant gap that affects daily user interaction with the system.