# CyFir Deployment Safety Notes

## ⚠️ Critical Safety Considerations

### 1. What We Did NOT Change (For Safety)

#### Tool Names Remain Unchanged
```yaml
# These remain as-is to prevent artifact breakage
- VelociraptorWindows
- VelociraptorLinux  
- VelociraptorCollector
```
**Why**: Changing these would break all existing artifacts and hunts.

#### Service Names Support Both
- Old: `velociraptor_client`, `Velociraptor`
- New: `cyfir`, `CyFir`
- **Both work simultaneously**

#### Import Paths Are Internal
- Changed: Source code imports
- Unchanged: User artifacts, VQL queries
- **No impact on custom content**

### 2. Safe Testing Approach

#### Start with Read-Only Testing
```bash
# Test configuration loading only
./cyfir config validate

# Test GUI without making changes
./cyfir --config server.config.yaml gui --read_only

# Query mode only
./cyfir query "SELECT * FROM info()"
```

#### Test in Isolated Environment
1. Spin up separate test server
2. Use test client pool only
3. Run for minimum 1 week
4. Monitor all logs

### 3. Gradual Migration Path

#### Week 1: Shadow Mode
- Install as `cyfir` alongside `velociraptor`
- Keep old binary running
- Test new binary in parallel

#### Week 2: Soft Switch
- Update symlinks
- Keep both services registered
- Monitor closely

#### Week 3: Primary Switch
- Make CyFir primary
- Keep Velociraptor as backup
- Ready to rollback

#### Week 4: Cleanup
- Remove old binary
- Update documentation
- Celebrate! 🎉

### 4. What to Monitor

#### Key Metrics
```sql
-- Client connectivity
SELECT count(*) FROM clients() 
WHERE last_seen_at > now() - 600

-- Active hunts
SELECT hunt_id, state, client_count 
FROM hunts() 
WHERE state = 'RUNNING'

-- Recent errors
SELECT client_id, timestamp, message 
FROM client_logs() 
WHERE level = 'ERROR' 
AND timestamp > now() - 3600
```

#### Log Monitoring
```bash
# Watch for connection errors
tail -f /var/log/cyfir/* | grep -E "ERROR|WARN|connection"

# Monitor service status
watch -n 5 'systemctl status cyfir velociraptor'

# Check client versions
cyfir query "SELECT version FROM clients() GROUP BY version"
```

### 5. Emergency Procedures

#### If Clients Won't Connect
```bash
# 1. Check both configs work
cyfir --config /etc/velociraptor/server.config.yaml config validate

# 2. Verify certificates
openssl x509 -in server.pem -text -noout

# 3. Check firewall rules
iptables -L -n | grep 8000

# 4. Rollback if needed
systemctl stop cyfir
systemctl start velociraptor
```

#### If GUI Shows Errors
1. Clear browser cache
2. Check for mixed content warnings
3. Verify GUI certificate
4. Use incognito mode for testing

#### If Artifacts Fail
- Check tool availability
- Verify permissions unchanged
- Review artifact parameters
- Test with simple artifact first

### 6. DO NOT Do These

#### ❌ Avoid These Actions
1. **Don't change tool names** in artifacts
2. **Don't remove old service** immediately  
3. **Don't update all clients** at once
4. **Don't skip testing** phase
5. **Don't ignore user feedback**

#### ❌ Risky Changes to Avoid
```yaml
# DON'T change these in artifacts:
tools:
  - name: VelociraptorWindows  # Keep as-is!

# DON'T remove old paths immediately:
writeback_linux: /etc/velociraptor.writeback.yaml  # Still works!

# DON'T force new env vars only:
# Keep supporting VELOCIRAPTOR_CONFIG
```

### 7. Safe Communication

#### Internal Communication Template
```
Subject: [INFO] Endpoint System Update - Testing Phase

Hi team,

We're testing an update to our endpoint system (Velociraptor → CyFir).
This is primarily a visual/branding change.

Current Status: Testing on isolated systems
Impact: None to production
Action Required: None

If you notice any issues with endpoint visibility, please report immediately.

Thanks,
Security Team
```

#### User Communication (If Needed)
```
The security team has updated our endpoint monitoring system.
You may see "CyFir" instead of "Velociraptor" in some places.
No action is required from you.
```

### 8. Validation Checklist

Before going to production, verify:

- [ ] Both binaries work: `velociraptor` and `cyfir`
- [ ] Both env vars work: `VELOCIRAPTOR_CONFIG` and `CYFIR_CONFIG`
- [ ] GUI shows CyFir branding
- [ ] Artifacts still execute successfully
- [ ] Clients maintain connection
- [ ] No errors in logs
- [ ] Performance metrics stable
- [ ] Rollback procedure tested

### 9. Support Resources

#### Documentation
- User Guide: `CYFIR_USER_GUIDE.md`
- Migration Guide: `CYFIR_MIGRATION_GUIDE.md`
- Test Results: `CYFIR_TEST_RESULTS.md`

#### Getting Help
1. Check deployment checklist first
2. Review test results document
3. Search logs for specific errors
4. Contact vendor support if needed

### 10. Final Safety Reminder

> **Golden Rule**: If anything seems wrong, stop and investigate. 
> It's better to delay deployment than to break production.

Remember:
- This is a **brand change**, not a functionality change
- **Compatibility** is preserved everywhere
- **Gradual migration** is the safest approach
- **Monitoring** is your best friend
- **Rollback** is always an option

---

*Safety first. Test thoroughly. Deploy gradually. Monitor constantly.*