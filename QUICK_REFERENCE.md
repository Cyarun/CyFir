# CyFir Rebranding - Quick Reference

## Current State (Phase 2B-2 Complete)
- ✅ Environment variables: Both VELOCIRAPTOR_* and CYFIR_* work
- ✅ Binaries: Both velociraptor and cyfir exist and work
- ✅ 11 safe strings updated
- 📊 318 Velociraptor references remain

## Key Commands

### Testing
```bash
# Run test suite
./run_phase2_tests.sh

# Test environment variables
CYFIR_CONFIG=test_config.yaml ./cyfir version
VELOCIRAPTOR_CONFIG=test_config.yaml ./cyfir version

# Test config generation
./cyfir config generate --interactive
```

### Building
```bash
# Quick build
go build -tags "server_vql" -o cyfir ./bin/

# Full build (requires b0x.yaml files)
make auto
```

### Git Checkpoints
```bash
# View tags
git tag -l phase*

# Rollback if needed
git reset --hard phase2b1-complete  # Before Phase 2B-2
git reset --hard phase2a-complete    # Before string changes
```

## Important Files
- `config/branding.go` - Branding configuration (KEY DISCOVERY!)
- `PHASE2_FINAL_SUMMARY.md` - Detailed summary
- `PHASE2B2_DEFER.md` - Strings we didn't change and why
- `run_phase2_tests.sh` - Test automation

## Sensitive Areas (DO NOT CHANGE YET)
- Registry paths: `HKLM\SOFTWARE\Velocidex\Velociraptor`
- Service names in code
- Certificate CN/SAN fields
- ETW provider names
- Protocol identifiers

## Next Strategic Decision
Choose approach for remaining 318 references:
1. **Branding Config** (recommended) - Use existing system
2. **Manual Updates** - Continue current approach
3. **Hybrid** - Mix of both

## Safe to Change
- Comments
- Help text
- User-facing messages
- Log messages (verify not parsed)
- Documentation

## Contact for Questions
- Review git log for detailed history
- Check PHASE2_*.md files for context
- All changes are tagged and documented