# GUI Rebranding Plan

## Files to Update:

### 1. High Priority (User-Visible):
- `gui/velociraptor/src/index.html` - Page title
- `gui/velociraptor/package.json` - Package name
- `gui/velociraptor/src/components/welcome/login.jsx` - Login screen
- `gui/velociraptor/src/components/i8n/en.jsx` - English translations

### 2. Internationalization (All Languages):
- Update all language files in `src/components/i8n/`
- Languages: en, fr, de, es, por, vi, jp
- Both manual and automated translations

### 3. Safe Update Strategy:
```javascript
// Add configuration option
const BRANDING = {
  name: process.env.REACT_APP_BRAND_NAME || "CyFir",
  fullName: "CyFir - Cyber Forensics & IR Platform",
  company: "CynorSense Solutions Pvt. Ltd.",
  legacy: "Velociraptor" // For compatibility
};
```

### 4. Phased Approach:
1. Add branding configuration
2. Update English UI first
3. Update other languages
4. Update logos/icons
5. Update help documentation