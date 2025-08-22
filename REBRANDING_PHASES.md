# CyFir Rebranding - Phased Approach

## Phase 1: Core Rebranding (Current - 60% Complete)
**Goal**: Change visible branding while maintaining functionality

### Safe Changes ✅
- Binary names (velociraptor → cyfir)
- UI text and labels
- Documentation strings
- Copyright notices
- Package names
- Log messages
- Error messages for users

### DO NOT CHANGE ❌
- Function names
- API endpoints
- Protocol definitions
- File formats
- VQL language name (yet)
- Internal constants
- Test assertions

## Phase 2: Infrastructure Migration (Future)
**Goal**: Move to CynorSense infrastructure

### 2.1 Artifact Repository
- Create https://github.com/CyFir-Artifacts
- Implement artifact sync from GitHub
- Update artifact loader to pull from GitHub
- Maintain embedded artifacts for offline use

### 2.2 Documentation Migration
- Set up GitHub Wiki
- Migrate all docs from cyfir.cynorsense.com
- Update all documentation links
- Create API documentation

### 2.3 VQL → CYF Transition
**CAREFUL**: This is a breaking change!

1. **Preparation**:
   - Create compatibility layer
   - Update parser to accept both "VQL" and "CYF"
   - Add deprecation warnings

2. **Migration Tools**:
   ```go
   // Artifact converter
   type QueryConverter struct {
       // Convert VQL to CYF in artifacts
       ConvertArtifact(yaml []byte) ([]byte, error)
       // Validate converted queries
       ValidateQuery(query string) error
   }
   ```

3. **Staged Rollout**:
   - v1.1: Accept both VQL and CYF
   - v1.2: Warn on VQL usage
   - v2.0: CYF only

## Phase 3: Complete Ecosystem Migration
- Update all external tools
- Migrate community artifacts
- Update training materials
- Deprecate old endpoints

## Current Safe Replacements

### String Replacements (SAFE)
```bash
# User-visible strings only
"Velociraptor server" → "CyFir server"
"Velociraptor client" → "CyFir client"
"Start Velociraptor" → "Start CyFir"
"Velociraptor is running" → "CyFir is running"
```

### Technical Terms (KEEP AS-IS)
```bash
# Do not change these:
- VQL (Velociraptor Query Language)
- velocidex.com in package imports
- API endpoint names
- Function names like GetVelociraptorConfig()
- Proto field names
```

## Implementation Checklist

### Phase 1 Tasks (Now)
- [x] Update binary names
- [x] Update UI strings
- [ ] Update error messages safely
- [ ] Update documentation strings
- [ ] Create compatibility wrappers
- [ ] Test all changes

### Phase 2 Tasks (Later)
- [ ] Design CYF language spec
- [ ] Create GitHub artifact repo
- [ ] Set up GitHub Wiki
- [ ] Build migration tools
- [ ] Create compatibility layer

### Phase 3 Tasks (Future)
- [ ] Full VQL → CYF migration
- [ ] Deprecate old APIs
- [ ] Update all tooling
- [ ] Community migration

## Risk Mitigation

1. **Testing Strategy**
   - Run full test suite after each change
   - Test client-server compatibility
   - Test artifact execution
   - Test API compatibility

2. **Rollback Plan**
   - Git tags at each milestone
   - Compatibility mode flags
   - Dual binary support

3. **Communication**
   - Migration guide for users
   - API compatibility notes
   - Artifact conversion tools