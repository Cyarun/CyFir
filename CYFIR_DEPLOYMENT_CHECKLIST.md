# CyFir Safe Deployment Checklist

## Pre-Deployment Safety Checks

### 1. Code Review ✓
- [ ] Review all changes in git log
- [ ] Verify no sensitive data in commits
- [ ] Check for any hardcoded credentials
- [ ] Ensure all TODOs are addressed

### 2. Testing Environment
- [ ] Set up isolated test environment
- [ ] Deploy to staging server first
- [ ] Test client-server communication
- [ ] Verify artifact collection works
- [ ] Test tool downloads

### 3. Backup Current Production
- [ ] Backup server configuration
- [ ] Backup datastore
- [ ] Document current service names
- [ ] Save current binary version
- [ ] Export critical artifacts

## Deployment Steps

### Phase 1: Pilot Deployment (Week 1-2)

#### 1.1 Small Test Group
```bash
# Deploy to 5-10 test clients first
- [ ] Select non-critical test machines
- [ ] Deploy new binary as 'cyfir'
- [ ] Keep velociraptor symlink
- [ ] Monitor for 48 hours
- [ ] Check logs for errors
```

#### 1.2 Validation Points
- [ ] Clients connect successfully
- [ ] Hunts execute properly
- [ ] Artifacts collect data
- [ ] GUI displays correctly
- [ ] No performance degradation

### Phase 2: Gradual Rollout (Week 3-4)

#### 2.1 Department by Department
- [ ] IT/Security team first (can troubleshoot)
- [ ] Non-critical departments next
- [ ] Critical systems last
- [ ] Document any issues

#### 2.2 Monitoring
```bash
# Monitor key metrics
- Client connection rates
- Collection success rates
- Error logs
- Performance metrics
- User feedback
```

### Phase 3: Full Production (Week 5+)

#### 3.1 Final Deployment
- [ ] Update all remaining clients
- [ ] Update documentation wiki
- [ ] Train support staff
- [ ] Update monitoring dashboards

## Rollback Plan

### Immediate Rollback Triggers
- [ ] Clients cannot connect
- [ ] Data collection fails
- [ ] Critical errors in logs
- [ ] Performance degradation >20%

### Rollback Steps
```bash
# 1. Stop services
systemctl stop cyfir

# 2. Restore old binary
cp /backup/velociraptor /usr/local/bin/
ln -sf velociraptor cyfir

# 3. Restore configuration if needed
cp /backup/server.config.yaml /etc/velociraptor/

# 4. Restart services
systemctl start velociraptor

# 5. Verify clients reconnect
```

## Communication Plan

### 1. Pre-Deployment (1 week before)
```markdown
Subject: Upcoming CyFir Update - No Action Required

Team,

We'll be updating our endpoint monitoring system from Velociraptor 
to CyFir over the next few weeks. This is primarily a branding 
change with no functional impact.

What to expect:
- New login screen saying "CyFir"
- Same functionality as before
- No action required from you

IT Security Team
```

### 2. During Deployment
- [ ] Status updates every 2 days
- [ ] Issues tracking spreadsheet
- [ ] Dedicated Slack channel
- [ ] Daily standup meetings

### 3. Post-Deployment
- [ ] Success announcement
- [ ] Feedback survey
- [ ] Lessons learned document

## Safety Configuration

### 1. Dual-Mode Operation
```yaml
# Keep both names during transition
windows_installer:
  service_name: CyFir
  # Old service remains for compatibility
  
Client:
  # Both binaries work
  # /usr/local/bin/velociraptor -> cyfir
```

### 2. Environment Variables
```bash
# Set both during transition
export VELOCIRAPTOR_CONFIG=/etc/velociraptor/server.config.yaml
export CYFIR_CONFIG=/etc/velociraptor/server.config.yaml
```

### 3. Monitoring Both Names
```bash
# Monitor logs for both
tail -f /var/log/velociraptor/* /var/log/cyfir/*

# Check both service names
systemctl status velociraptor cyfir
```

## Validation Scripts

### 1. Pre-Deployment Validation
```bash
#!/bin/bash
# validate_deployment.sh

echo "=== CyFir Pre-Deployment Validation ==="

# Check binary
if ./cyfir version | grep -q "name: cyfir"; then
    echo "✓ Binary shows correct name"
else
    echo "✗ Binary name incorrect"
    exit 1
fi

# Check config loading
export VELOCIRAPTOR_CONFIG=./server.config.yaml
if ./cyfir config validate; then
    echo "✓ Old env var works"
else
    echo "✗ Old env var broken"
    exit 1
fi

# Check symlink
if [ -L "./velociraptor" ]; then
    echo "✓ Compatibility symlink exists"
else
    echo "✗ Missing compatibility symlink"
    exit 1
fi

echo "=== All checks passed ==="
```

### 2. Post-Deployment Health Check
```bash
#!/bin/bash
# health_check.sh

# Check service status
if systemctl is-active --quiet cyfir || systemctl is-active --quiet velociraptor; then
    echo "✓ Service is running"
else
    echo "✗ Service is not running"
fi

# Check client connections
CLIENTS=$(cyfir query "SELECT count(*) FROM clients() WHERE last_seen_at > now() - 3600" | grep -o '[0-9]\+')
echo "✓ Active clients in last hour: $CLIENTS"

# Check for errors
ERRORS=$(journalctl -u cyfir -u velociraptor --since "1 hour ago" | grep -c ERROR)
if [ $ERRORS -eq 0 ]; then
    echo "✓ No errors in last hour"
else
    echo "⚠ Found $ERRORS errors in logs"
fi
```

## Risk Mitigation

### 1. Known Risks
- **Tool Downloads**: Keep old tool names
- **Service Names**: Support both old and new
- **API Integrations**: May need updates
- **Automation Scripts**: Might reference old name

### 2. Mitigation Strategies
- Maintain dual support for 6 months
- Clear communication about changes
- Provide migration scripts
- Keep detailed logs

### 3. Emergency Contacts
```
Primary: Security Team Lead
Secondary: Infrastructure Manager
Escalation: CTO
Vendor Support: support@cynorsense.com
```

## Success Criteria

### Week 1
- [ ] 100% test clients connected
- [ ] Zero critical errors
- [ ] All artifacts working

### Week 2
- [ ] 50% production migrated
- [ ] <5 support tickets
- [ ] No rollbacks needed

### Week 4
- [ ] 100% migration complete
- [ ] User satisfaction >90%
- [ ] No performance impact

## Post-Deployment Tasks

### 1. Documentation Updates
- [ ] Update wiki/confluence
- [ ] Update runbooks
- [ ] Update training materials
- [ ] Update architecture diagrams

### 2. Cleanup (After 3 months)
- [ ] Remove old log directories
- [ ] Remove old service entries
- [ ] Update all scripts to use cyfir
- [ ] Archive old documentation

### 3. Lessons Learned
- [ ] Document what went well
- [ ] Document issues encountered
- [ ] Update deployment process
- [ ] Share with community

---

## Sign-Off

- [ ] Security Team Lead: ________________
- [ ] Infrastructure Manager: ________________  
- [ ] CTO Approval: ________________
- [ ] Deployment Date: ________________

*Remember: Safety first. When in doubt, delay and test more.*