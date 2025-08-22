# CyFir Migration Guide

## Migrating from Velociraptor to CyFir

This guide helps existing Velociraptor deployments migrate to CyFir by CynorSense Solutions.

### Overview

CyFir is a rebranded version of Velociraptor, maintaining full compatibility while introducing the new CynorSense Solutions branding. The migration process is designed to be gradual and non-disruptive.

### Migration Phases

#### Phase 1: Preparation (Current)
- [ ] Backup existing Velociraptor configuration
- [ ] Document current deployment architecture
- [ ] Test CyFir binary in isolated environment
- [ ] Review custom artifacts and configurations

#### Phase 2: Compatibility Mode
- [ ] Deploy CyFir server alongside existing Velociraptor
- [ ] Use service compatibility scripts for dual operation
- [ ] Test client connectivity with both names
- [ ] Verify artifact compatibility

#### Phase 3: Gradual Transition
- [ ] Update clients using built-in upgrade artifacts
- [ ] Migrate server configurations
- [ ] Update custom artifacts to CyFir branding
- [ ] Transition monitoring and alerting

#### Phase 4: Completion
- [ ] Decommission old Velociraptor services
- [ ] Update documentation and procedures
- [ ] Remove compatibility layers
- [ ] Full CyFir deployment

### Technical Details

#### Binary Compatibility

CyFir maintains binary compatibility with Velociraptor:

```bash
# Old command
velociraptor server --config server.config.yaml

# New command (both work)
cyfir server --config server.config.yaml
velociraptor server --config server.config.yaml  # Compatibility alias
```

#### Service Names

The service installation scripts support both names:

```bash
# Systemd services
systemctl status velociraptor_server  # Old
systemctl status cyfir_server         # New

# Both names work during transition
```

#### Configuration Files

No changes required to configuration files. CyFir uses the same format:

```yaml
# Existing server.config.yaml works unchanged
Client:
  server_urls:
    - https://cyfir.cynorsense.com:8000/
```

#### Environment Variables

All environment variables remain compatible:

```bash
# These remain unchanged
VELOCIRAPTOR_CONFIG=/etc/velociraptor/client.config.yaml
VELOCIRAPTOR_API_CONFIG=/etc/velociraptor/api.config.yaml
```

### Client Migration

#### Option 1: In-Place Upgrade

Use the Admin.Client.Upgrade artifacts:

```vql
SELECT * FROM Artifact.Admin.Client.Upgrade.Windows(
  SleepDuration=600
)
```

#### Option 2: Fresh Deployment

Deploy new CyFir clients alongside existing ones:

1. Generate new client configs with CyFir branding
2. Deploy using existing deployment tools
3. Gradually decommission old clients

#### Option 3: Repackaging

Create new MSI/DEB/RPM packages:

```bash
# Create new packages with CyFir branding
cyfir config client --name CyFirClient > client.config.yaml
cyfir rpm client --config client.config.yaml
cyfir deb client --config client.config.yaml
```

### Server Migration

#### Step 1: Parallel Deployment

1. Install CyFir server on same or new hardware
2. Copy existing server configuration
3. Update server URLs in configuration
4. Start CyFir server on different port initially

#### Step 2: Data Migration

```bash
# Copy filestore data
rsync -av /var/lib/velociraptor/ /var/lib/cyfir/

# Update permissions
chown -R cyfir:cyfir /var/lib/cyfir/
```

#### Step 3: DNS Transition

1. Create new DNS entry: cyfir.cynorsense.com
2. Test client connectivity to new endpoint
3. Update client configurations gradually
4. Deprecate old DNS entry

### Artifact Migration

Custom artifacts need minimal updates:

```yaml
# Old artifact
name: Custom.Velociraptor.Collection
description: Velociraptor custom artifact

# Updated artifact  
name: Custom.CyFir.Collection
description: CyFir custom artifact
```

### API Integration

API endpoints remain compatible:

```python
# Python API client works unchanged
import pyvelociraptor
from pyvelociraptor import api_pb2

# Connection works with new server
config = pyvelociraptor.LoadConfigFile("api.config.yaml")
```

### Monitoring Integration

Update monitoring configurations:

```yaml
# Prometheus metrics
- job_name: 'cyfir'
  static_configs:
    - targets: ['cyfir.cynorsense.com:8003']
      labels:
        instance: 'cyfir-prod'
```

### Common Issues and Solutions

#### Issue: Clients Can't Connect
- Verify DNS resolution for new domain
- Check firewall rules for CyFir ports
- Ensure certificates are valid for new domain

#### Issue: Service Name Conflicts
- Use compatibility scripts during transition
- Stop old service before starting new one
- Check for port conflicts

#### Issue: Custom Tools Break
- Update tool references in artifacts
- Reupload tools with new names if needed
- Test tools in staging environment

### Rollback Plan

If issues arise during migration:

1. Stop CyFir services
2. Restart Velociraptor services
3. Revert DNS changes
4. Investigate issues in test environment

### Support

For migration assistance:
- Documentation: https://cyfir.cynorsense.com/docs
- Support: support@cynorsense.com
- GitHub: https://github.com/Cyarun/CyFir

### Checklist

- [ ] Backup configurations and data
- [ ] Test in staging environment
- [ ] Create migration timeline
- [ ] Notify stakeholders
- [ ] Update documentation
- [ ] Train team on new branding
- [ ] Monitor migration progress
- [ ] Validate functionality post-migration
- [ ] Decommission old infrastructure
- [ ] Celebrate successful migration! 🎉