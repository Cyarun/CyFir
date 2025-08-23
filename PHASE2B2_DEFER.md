# Phase 2B-2 Deferred Changes

## Registry Path Reference
- File: `tools/survey/server.go:96`
- Text: `By default we use HKLM\SOFTWARE\Velocidex\Velociraptor`
- **Reason for deferring**: This describes the actual registry path used. Changing this would require:
  1. Changing the actual registry path in code
  2. Migration logic for existing installations
  3. Extensive testing on Windows
- **Recommendation**: Address in Phase 3 with proper migration plan