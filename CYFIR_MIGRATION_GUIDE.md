# CyFir Migration Guide

## Welcome to CyFir!

Velociraptor is now CyFir, powered by CynorSense Solutions. This guide helps you transition smoothly to CyFir while maintaining your existing deployments.

## What's Changed?

### Visual Changes
- Application now displays "CyFir" branding
- Banner shows CynorSense Solutions
- GUI login shows "CyFir Login"
- Help text references CyFir

### What Hasn't Changed
- ✅ All functionality remains identical
- ✅ Existing configurations still work
- ✅ Client-server compatibility maintained
- ✅ No data migration required

## Migration Options

### Option 1: No Action Required (Recommended)
Your existing Velociraptor installation will continue to work perfectly. The CyFir update maintains full backward compatibility.

### Option 2: Gradual Migration
Update at your own pace:

1. **Update Environment Variables** (Optional)
   ```bash
   # Old (still works)
   export VELOCIRAPTOR_CONFIG=/path/to/config.yaml
   
   # New (recommended)
   export CYFIR_CONFIG=/path/to/config.yaml
   ```

2. **Update Binary Names** (Optional)
   ```bash
   # Both commands work identically
   velociraptor version
   cyfir version
   ```

3. **Update Scripts** (When convenient)
   Replace `velociraptor` with `cyfir` in your scripts

## Quick Start Commands

### Check Your Version
```bash
# Either command works
velociraptor version
cyfir version
```

### Start Server
```bash
# Using old name (works)
velociraptor --config server.yaml frontend

# Using new name (works)
cyfir --config server.yaml frontend
```

### Generate New Config
```bash
cyfir config generate > server.yaml
```

## Environment Variables

Both old and new variables work. New variables take precedence if both exist.

| Old Variable | New Variable | Purpose |
|-------------|--------------|---------|
| VELOCIRAPTOR_CONFIG | CYFIR_CONFIG | Configuration file path |
| VELOCIRAPTOR_LITERAL_CONFIG | CYFIR_LITERAL_CONFIG | Inline configuration |
| VELOCIRAPTOR_API_CONFIG | CYFIR_API_CONFIG | API configuration path |

## Service Management

### Linux (systemd)
Service names remain unchanged for compatibility:
```bash
sudo systemctl status velociraptor
sudo systemctl restart velociraptor
```

### Windows
Service names remain unchanged:
```powershell
Get-Service Velociraptor
Restart-Service Velociraptor
```

## Frequently Asked Questions

### Q: Do I need to update immediately?
**A**: No. Your existing installation continues to work without any changes.

### Q: Will my clients still connect?
**A**: Yes. Client-server compatibility is fully maintained.

### Q: Do I need to regenerate certificates?
**A**: No. Existing certificates continue to work.

### Q: Can I mix old and new clients?
**A**: Yes. Old and new clients work together seamlessly.

### Q: What about my existing data?
**A**: All data formats remain unchanged. No migration needed.

## Troubleshooting

### Issue: Command not found
```bash
# If 'cyfir' command not found, use 'velociraptor'
velociraptor --help

# Or create an alias
alias cyfir=velociraptor
```

### Issue: Environment variable not working
```bash
# Check which variable is set
echo $VELOCIRAPTOR_CONFIG
echo $CYFIR_CONFIG

# Use either one
export CYFIR_CONFIG=$VELOCIRAPTOR_CONFIG
```

## Best Practices

1. **Test First**: Try CyFir in a test environment
2. **Update Documentation**: Update your internal docs to reference CyFir
3. **Gradual Rollout**: Update components as maintenance windows allow
4. **Keep Compatibility**: No need to update everything at once

## Support

### Documentation
- Official Docs: https://cyfir.cynorsense.com
- GitHub: https://github.com/Cyarun/CyFir

### Contact
- Support: support@cynorsense.com
- Website: https://cynorsense.com

## Summary

The transition from Velociraptor to CyFir is designed to be seamless:
- ✅ No breaking changes
- ✅ Full backward compatibility
- ✅ Update at your own pace
- ✅ Zero downtime migration

Welcome to CyFir - the same powerful platform you trust, with a new name!

---
*CyFir - Cyber Forensics & IR Platform by CynorSense Solutions*