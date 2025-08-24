# CyFir v2.0 Migration Roadmap

## Overview

This document outlines the planned breaking changes for CyFir v2.0, which will complete the rebranding by updating internal identifiers that couldn't be changed in v1.x due to compatibility requirements.

## Timeline

- **v1.x (Current)**: Full compatibility mode - no breaking changes
- **v1.9**: Deprecation warnings added
- **v2.0**: Breaking changes with migration tools
- **v2.1+**: Legacy support removed

## Planned Breaking Changes

### 1. Tool Name Updates

#### Current State (v1.x)
```yaml
tools:
  - name: VelociraptorWindows
  - name: VelociraptorLinux
  - name: VelociraptorCollector
  - name: VelociraptorDebian
  - name: VelociraptorRedHat
```

#### Target State (v2.0)
```yaml
tools:
  - name: CyFirWindows
  - name: CyFirLinux
  - name: CyFirCollector
  - name: CyFirDebian
  - name: CyFirRedHat
```

#### Migration Strategy
1. **Dual Registration Period** (v1.9)
   - Register both old and new tool names
   - Emit deprecation warnings for old names
   - Auto-redirect old names to new

2. **Artifact Migration Tool**
   ```bash
   cyfir migrate artifacts --update-tools
   ```
   - Automatically updates all artifact files
   - Creates backup of original files
   - Validates artifact syntax after update

### 2. Service Name Standardization

#### Current State (v1.x)
- Windows: "Velociraptor" or "CyFir"
- Linux: "velociraptor_client" or "cyfir"
- macOS: "com.rychlabs.velociraptor" or "com.cynorsense.cyfir"

#### Target State (v2.0)
- Windows: "CyFir"
- Linux: "cyfir"
- macOS: "com.cynorsense.cyfir"

#### Migration Tool
```bash
cyfir migrate service --preserve-old
```

Features:
- Detects existing service
- Creates new service with data migration
- Optionally removes old service
- Handles service dependencies

### 3. Environment Variable Cleanup

#### Deprecated in v2.0
- `VELOCIRAPTOR_CONFIG`
- `VELOCIRAPTOR_LITERAL_CONFIG`
- `VELOCIRAPTOR_API_CONFIG`

#### Migration Path
1. **v1.9 Warning Phase**
   ```
   WARNING: VELOCIRAPTOR_CONFIG is deprecated. Please use CYFIR_CONFIG.
   ```

2. **v2.0 Enforcement**
   - Old variables ignored
   - Clear error message with migration instructions

### 4. Internal Constants

#### Updates Required
```go
// v1.x
VELOCIRAPTOR_WRITEBACK = "/etc/velociraptor.writeback.yaml"
VELOCIRAPTOR_BUFFER = "Velociraptor_Buffer.bin"

// v2.0
CYFIR_WRITEBACK = "/etc/cyfir.writeback.yaml"
CYFIR_BUFFER = "CyFir_Buffer.bin"
```

#### Compatibility Mode
- Check both locations during transition
- Automatic migration of existing files
- Preserve file permissions and ownership

### 5. API Endpoints

#### Current Compatibility Endpoints
```
/api/v1/GetVelociraptorVersion
/api/v1/VelociraptorConfig
```

#### v2.0 Changes
- Old endpoints return 301 redirects
- New endpoints: `/api/v1/GetCyFirVersion`
- Deprecation headers in responses

## Migration Tools Suite

### 1. Pre-Migration Checker
```bash
cyfir migrate check
```
Output:
```
CyFir v2.0 Migration Readiness Check
====================================
✅ Configuration compatible
⚠️  3 artifacts use old tool names
⚠️  Service uses old name: velociraptor_client
✅ No custom API integrations detected
❌ Environment using VELOCIRAPTOR_CONFIG

Ready for migration: NO
Run 'cyfir migrate prepare' for detailed instructions
```

### 2. Migration Preparation
```bash
cyfir migrate prepare --output migration-plan.md
```
Generates:
- Detailed migration steps
- Backup instructions
- Rollback procedures
- Testing checklist

### 3. Automated Migration
```bash
cyfir migrate apply --backup-dir /backup/v1
```
Actions:
- Creates full backup
- Updates configurations
- Migrates artifacts
- Updates service registrations
- Validates deployment

### 4. Rollback Tool
```bash
cyfir migrate rollback --backup-dir /backup/v1
```

## Communication Plan

### 6 Months Before v2.0
- Blog post announcing v2.0 plans
- Migration guide published
- Beta program launched

### 3 Months Before v2.0
- Deprecation warnings in v1.9
- Migration tools released
- Webinar for administrators

### 1 Month Before v2.0
- Final reminder emails
- Updated documentation
- Support team training

### v2.0 Release
- Comprehensive release notes
- Migration success stories
- Extended support period

## Support Strategy

### Transition Period (6 months)
- v1.x continues to receive security updates
- Migration support included
- Community forums for assistance

### Long-term Support
- v1.x LTS version for 1 year
- Security patches only
- Clear EOL communication

## Risk Mitigation

### 1. Automated Testing
- Migration tool test suite
- Compatibility validators
- Performance benchmarks

### 2. Beta Program
- Early access for key users
- Feedback incorporation
- Real-world testing

### 3. Rollback Capability
- One-command rollback
- Data integrity preserved
- Service continuity maintained

## Success Metrics

### Target Goals
- 80% migration within 6 months
- <1% failed migrations
- Zero data loss incidents
- 95% user satisfaction

### Monitoring
- Telemetry on version adoption
- Support ticket tracking
- Community sentiment analysis
- Performance metrics

## Conclusion

The v2.0 migration completes the CyFir rebranding while providing users with:
- Clear migration path
- Comprehensive tooling
- Safety mechanisms
- Extended support

This measured approach ensures successful adoption while minimizing disruption to production environments.

---

**Status**: PLANNED
**Target Release**: TBD (6+ months after v1.x adoption)
**Risk Level**: MEDIUM (with mitigation)
**User Impact**: CONTROLLED