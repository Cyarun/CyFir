# CyFir Branding Assets

This directory contains all branding assets for CyFir by CynorSense Solutions.

## Directory Structure

```
branding/
├── logos/          # Logo files in various formats
├── icons/          # Application icons
├── banners/        # Marketing banners and headers
└── guidelines/     # Brand usage guidelines
```

## Logo Specifications

### Primary Logo
- **File**: `cyfir_logo.svg`
- **Colors**: 
  - Primary: #00AA00 (Green)
  - Secondary: #000000 (Black)
  - Background: #FFFFFF (White)

### Logo Variations
1. **Full Logo**: CyFir with tagline
2. **Icon Only**: C symbol
3. **Wordmark**: CyFir text only
4. **Monochrome**: Black/white versions

## Color Palette

### Primary Colors
- **CyFir Green**: #00AA00
- **Deep Black**: #000000
- **Pure White**: #FFFFFF

### Secondary Colors
- **Dark Green**: #008773
- **Light Green**: #00911E
- **Gray**: #666666

## Typography

### Primary Font
- **Headers**: Gotham, "Helvetica Neue", sans-serif
- **Body**: -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif
- **Code**: "Courier New", Courier, monospace

## Icon Sizes

### Application Icons
- Windows: 16x16, 32x32, 48x48, 256x256
- macOS: 16x16, 32x32, 64x64, 128x128, 256x256, 512x512, 1024x1024
- Linux: 16x16, 22x22, 24x24, 32x32, 48x48, 64x64, 128x128, 256x256

### Favicon
- 16x16, 32x32, 192x192 (PNG)
- favicon.ico (multi-resolution)

## Usage Guidelines

### Logo Usage
1. Maintain minimum clear space equal to the height of the 'C' in CyFir
2. Do not distort, rotate, or modify the logo
3. Use approved color variations only
4. Ensure sufficient contrast with background

### Text References
- **Company**: CynorSense Solutions Pvt. Ltd.
- **Product**: CyFir
- **Tagline**: "Cyber Forensics & Incident Response Platform"
- **Domain**: cyfir.cynorsense.com

### File Naming Convention
```
cyfir_[type]_[variant]_[size].[format]

Examples:
- cyfir_logo_primary_256x256.png
- cyfir_icon_mono_32x32.png
- cyfir_banner_web_1920x480.jpg
```

## Implementation

### Web Usage
```html
<!-- Logo -->
<img src="/assets/logos/cyfir_logo.svg" alt="CyFir" height="40">

<!-- Favicon -->
<link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
<link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
```

### Application Resources
```go
//go:embed branding/icons/cyfir_icon_256x256.png
var AppIcon []byte

//go:embed branding/logos/cyfir_logo.svg
var AppLogo []byte
```

## Brand Protection

### Trademark
- CyFir™ is a trademark of CynorSense Solutions Pvt. Ltd.
- All rights reserved

### Copyright Notice
```
Copyright © 2024 CynorSense Solutions Pvt. Ltd.
All rights reserved.
```

## Contact

For branding inquiries:
- Email: branding@cynorsense.com
- Website: https://cynorsense.com