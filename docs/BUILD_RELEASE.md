# CyFir Build and Release Guide

## Overview

This guide covers building and releasing CyFir binaries for all supported platforms.

## Prerequisites

### Development Environment

1. **Go 1.21+** - [Download](https://golang.org/dl/)
2. **Node.js 18+ LTS** - [Download](https://nodejs.org/)
3. **Make** - Build automation tool
4. **Git** - Version control

### Platform-Specific Requirements

#### Linux
```bash
# Ubuntu/Debian
sudo apt-get install build-essential mingw-w64 gcc-multilib

# Fedora/RHEL
sudo dnf install gcc gcc-c++ mingw64-gcc golang

# OpenSUSE
sudo zypper install gcc gcc-c++ mingw64-gcc golangci-lint
```

#### macOS
```bash
# Install Xcode Command Line Tools
xcode-select --install

# Install dependencies via Homebrew
brew install go node mingw-w64
```

#### Windows
- Install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/)
- Install [Git for Windows](https://git-scm.com/download/win)
- Install [Node.js](https://nodejs.org/)
- Install [Go](https://golang.org/dl/)

## Building CyFir

### Quick Build (Development)

```bash
# Clone the repository
git clone https://github.com/Cyarun/CyFir.git
cd CyFir

# Build GUI assets
cd gui/velociraptor
npm install
npm run build
cd ../..

# Build development binary with race detection
make

# Binary will be in output/velociraptor (or output/cyfir)
```

### Production Builds

#### Linux Binary
```bash
make linux
# Output: output/velociraptor_prod_linux_amd64
```

#### Windows Binary
```bash
make windows
# Output: output/velociraptor_prod_windows_amd64.exe
```

#### macOS Binary
```bash
make darwin
# Output: output/velociraptor_prod_darwin_amd64
```

#### All Platforms
```bash
make release
# Builds all platform binaries
```

### Build Options

```bash
# Build without CGO (static binary)
CGO_ENABLED=0 make linux

# Build with specific version
make linux VERSION=1.0.0-cyfir

# Build with custom flags
make EXTRA_GO_FLAGS="-ldflags='-s -w'" linux
```

## Creating Release Packages

### Linux Packages

#### Debian Package (.deb)
```bash
# Build client package
./output/velociraptor debian client --binary output/velociraptor_prod_linux_amd64

# Build server package  
./output/velociraptor debian server --binary output/velociraptor_prod_linux_amd64
```

#### RPM Package
```bash
# Build client package
./output/velociraptor rpm client --binary output/velociraptor_prod_linux_amd64

# Build server package
./output/velociraptor rpm server --binary output/velociraptor_prod_linux_amd64
```

### Windows Packages

#### MSI Installer
```bash
# Generate MSI for Windows
./output/velociraptor config client --format msi > cyfir_client.msi
```

#### Windows Service Installer
```powershell
# Install as Windows service
.\velociraptor.exe service install --config client.config.yaml
```

### macOS Package

```bash
# Create macOS app bundle
make darwin_app

# Create DMG installer
make darwin_dmg
```

## Release Process

### 1. Version Tagging

```bash
# Update version in constants
vim constants/constants.go

# Commit version change
git add -A
git commit -m "Release: CyFir v1.0.0"

# Create release tag
git tag -a v1.0.0 -m "CyFir v1.0.0 Release"
git push origin v1.0.0
```

### 2. Build Release Binaries

```bash
# Clean previous builds
make clean

# Build all release binaries
make release

# Verify builds
ls -la output/
```

### 3. Create Release Artifacts

```bash
# Create release directory
mkdir -p release/v1.0.0

# Copy binaries
cp output/velociraptor_prod_* release/v1.0.0/

# Create packages
cd release/v1.0.0
for binary in velociraptor_prod_*; do
    # Create DEB packages
    ../../output/velociraptor debian client --binary $binary --output cyfir-client-${binary##*_}.deb
    ../../output/velociraptor debian server --binary $binary --output cyfir-server-${binary##*_}.deb
    
    # Create RPM packages  
    ../../output/velociraptor rpm client --binary $binary --output cyfir-client-${binary##*_}.rpm
    ../../output/velociraptor rpm server --binary $binary --output cyfir-server-${binary##*_}.rpm
done

# Create checksums
sha256sum * > SHA256SUMS
```

### 4. GitHub Release

```bash
# Install GitHub CLI if needed
# https://cli.github.com/

# Create release
gh release create v1.0.0 \
  --title "CyFir v1.0.0" \
  --notes "Release notes here" \
  release/v1.0.0/*
```

## Continuous Integration

### GitHub Actions Workflow

The `.github/workflows/cyfir-ci.yml` workflow automatically:

1. Builds binaries for all platforms
2. Runs tests
3. Creates release artifacts
4. Publishes to GitHub Releases (on tags)

### Manual Trigger

```bash
# Trigger workflow manually
gh workflow run cyfir-ci.yml
```

## Build Verification

### Verify Binary

```bash
# Check version
./output/velociraptor version

# Verify binary info
file output/velociraptor
ldd output/velociraptor  # Linux only

# Test basic functionality
./output/velociraptor gui
```

### Run Tests

```bash
# Run all tests
make test

# Run specific package tests
go test -v ./services/...

# Run with race detection
make test_race
```

## Troubleshooting

### Common Build Issues

1. **Missing dependencies**
   ```bash
   go mod download
   cd gui/velociraptor && npm install
   ```

2. **CGO errors on Windows**
   - Ensure TDM-GCC is in PATH
   - Use `CGO_ENABLED=0` for static builds

3. **Node/npm errors**
   ```bash
   cd gui/velociraptor
   rm -rf node_modules package-lock.json
   npm install
   ```

4. **Cross-compilation issues**
   ```bash
   # Set proper environment
   GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \
   CC=x86_64-w64-mingw32-gcc \
   make windows
   ```

## Security Considerations

### Code Signing

#### Windows
```powershell
# Sign Windows binary
signtool sign /f cyfir.pfx /p password /tr http://timestamp.server cyfir.exe
```

#### macOS
```bash
# Sign macOS binary
codesign --deep --force --verify --verbose \
  --sign "Developer ID Application: CynorSense Solutions" \
  cyfir.app
```

### Package Signing

#### DEB Packages
```bash
# Sign DEB package
dpkg-sig --sign builder cyfir-client.deb
```

#### RPM Packages
```bash
# Sign RPM package
rpm --addsign cyfir-client.rpm
```

## Release Checklist

- [ ] Update version in `constants/constants.go`
- [ ] Update CHANGELOG.md
- [ ] Run full test suite
- [ ] Build all platform binaries
- [ ] Create installation packages
- [ ] Test installations on target platforms
- [ ] Generate checksums
- [ ] Sign binaries and packages
- [ ] Create GitHub release
- [ ] Update documentation
- [ ] Announce release

## Support

For build issues:
- GitHub Issues: https://github.com/Cyarun/CyFir/issues
- Documentation: https://cyfir.cynorsense.com/docs
- Email: support@cynorsense.com