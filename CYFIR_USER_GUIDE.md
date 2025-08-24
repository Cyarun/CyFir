# CyFir User Migration Guide

## Welcome to CyFir by CynorSense Solutions!

Velociraptor has been rebranded as **CyFir** - the same powerful digital forensics and incident response platform you know and trust, now with a fresh identity from CynorSense Solutions Pvt. Ltd.

## What's Changed?

### 🎨 Visual Changes
- **Application Name**: Velociraptor → CyFir
- **GUI Branding**: All user interfaces now display CyFir
- **Login Screen**: Shows "CyFir Login"
- **Documentation Links**: Point to cyfir.cynorsense.com
- **Theme Names**: "CyFir (light/dark)" themes available

### 🔧 Technical Changes
- **Binary Name**: `velociraptor` → `cyfir` (symlink provided for compatibility)
- **Environment Variables**: New `CYFIR_*` variables (old ones still work)
- **Configuration**: Service names updated to CyFir
- **Import Paths**: Using github.com/Cyarun/CyFir

## What Hasn't Changed?

### ✅ Full Compatibility Maintained
- **Existing Configs**: Your current configuration files work without changes
- **Service Names**: Existing Windows/Linux services continue to function
- **Tool Names**: Artifacts still use VelociraptorWindows, etc. (for now)
- **API Compatibility**: All APIs remain unchanged
- **Data Format**: No changes to data storage or formats

## Migration Steps

### For New Installations

1. **Download CyFir**
   ```bash
   # The binary is now named 'cyfir'
   wget https://github.com/Cyarun/CyFir/releases/latest/download/cyfir
   ```

2. **Generate Configuration**
   ```bash
   ./cyfir config generate > server.config.yaml
   ```

3. **Start Services**
   - Windows: Service name is "CyFir"
   - Linux: Service can be "cyfir" or "velociraptor_client"
   - macOS: Service is "com.cynorsense.cyfir"

### For Existing Installations

**No immediate action required!** Your current installation will continue to work.

#### Optional Updates

1. **Environment Variables** (Optional)
   ```bash
   # Old (still works)
   export VELOCIRAPTOR_CONFIG=/path/to/config.yaml
   
   # New (takes precedence)
   export CYFIR_CONFIG=/path/to/config.yaml
   ```

2. **Binary Name** (Optional)
   ```bash
   # A symlink is automatically created
   velociraptor → cyfir
   
   # Both commands work identically
   ./velociraptor version
   ./cyfir version
   ```

3. **Update Scripts** (Optional)
   Replace `velociraptor` with `cyfir` in your scripts for consistency.

## Configuration Changes

### Windows Installer
```yaml
windows_installer:
  service_name: CyFir
  install_path: $ProgramFiles\CyFir\CyFir.exe
  service_description: CyFir service
```

### macOS Installer
```yaml
darwin_installer:
  service_name: com.cynorsense.cyfir
  install_path: /usr/local/sbin/cyfir
```

### Documentation URL
```yaml
links:
- text: Documentation
  url: https://cyfir.cynorsense.com/
```

## Artifact Compatibility

All existing artifacts continue to work without modification:
- Tool names remain as `VelociraptorWindows`, etc.
- Artifact descriptions have been updated to reference CyFir
- No changes required to custom artifacts

## Troubleshooting

### Issue: Command not found
**Solution**: Use `cyfir` instead of `velociraptor`, or create a symlink:
```bash
ln -s cyfir velociraptor
```

### Issue: Service won't start
**Solution**: Check service name:
- Windows: Try both "CyFir" and "Velociraptor"
- Linux: Try both "cyfir" and "velociraptor_client"

### Issue: Environment variable not working
**Solution**: Set both variables during transition:
```bash
export VELOCIRAPTOR_CONFIG=/path/to/config.yaml
export CYFIR_CONFIG=/path/to/config.yaml
```

## Future Plans (v2.0)

In a future major release, we plan to:
- Update tool names (VelociraptorWindows → CyFirWindows)
- Complete service name migration
- Remove legacy compatibility layers

## Support

- **Documentation**: https://cyfir.cynorsense.com/
- **GitHub**: https://github.com/Cyarun/CyFir
- **Company**: https://cynorsense.com/

## FAQ

**Q: Do I need to update immediately?**
A: No, your existing installation will continue to work without any changes.

**Q: Will my artifacts break?**
A: No, all artifacts are fully compatible. Tool names remain unchanged.

**Q: Can I still use 'velociraptor' command?**
A: Yes, a symlink ensures the old command continues to work.

**Q: What about my automation scripts?**
A: They will continue to work. Update them to use 'cyfir' at your convenience.

**Q: Is this just a cosmetic change?**
A: The rebranding reflects our commitment to the product while maintaining full compatibility and the same powerful features you rely on.

---

*CyFir - Digging Deeper with CynorSense Solutions*