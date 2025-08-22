#!/bin/bash

echo "=== Checking CyFir/Velociraptor Build Dependencies ==="
echo

# Check Go version
echo "1. Checking Go version..."
if command -v go &> /dev/null; then
    GO_VERSION=$(go version | awk '{print $3}')
    echo "   ✓ Go installed: $GO_VERSION"
    
    # Extract version numbers
    REQUIRED="1.23.2"
    CURRENT=$(go version | awk '{print $3}' | sed 's/go//')
    echo "   Required: go$REQUIRED or higher"
    echo "   Current: go$CURRENT"
else
    echo "   ✗ Go is not installed!"
    exit 1
fi
echo

# Check for essential build tools
echo "2. Checking build tools..."
tools=("make" "gcc" "git")
for tool in "${tools[@]}"; do
    if command -v $tool &> /dev/null; then
        echo "   ✓ $tool is installed"
    else
        echo "   ✗ $tool is not installed"
    fi
done
echo

# Check for optional cross-compilation tools
echo "3. Checking optional cross-compilation tools..."
if command -v x86_64-w64-mingw32-gcc &> /dev/null; then
    echo "   ✓ MinGW cross compiler (64-bit) is installed"
else
    echo "   ℹ MinGW cross compiler (64-bit) not found (needed for Windows builds)"
fi

if command -v i686-w64-mingw32-gcc &> /dev/null; then
    echo "   ✓ MinGW cross compiler (32-bit) is installed"
else
    echo "   ℹ MinGW cross compiler (32-bit) not found (needed for Windows x86 builds)"
fi

if command -v musl-gcc &> /dev/null; then
    echo "   ✓ musl-gcc is installed"
else
    echo "   ℹ musl-gcc not found (needed for static Linux builds)"
fi
echo

# Check for Node.js (for GUI development)
echo "4. Checking GUI development tools..."
if command -v node &> /dev/null; then
    NODE_VERSION=$(node --version)
    echo "   ✓ Node.js installed: $NODE_VERSION"
else
    echo "   ℹ Node.js not found (needed for GUI development)"
fi

if command -v npm &> /dev/null; then
    NPM_VERSION=$(npm --version)
    echo "   ✓ npm installed: $NPM_VERSION"
else
    echo "   ℹ npm not found (needed for GUI development)"
fi
echo

# Check for testing tools
echo "5. Checking testing tools..."
if command -v staticcheck &> /dev/null; then
    echo "   ✓ staticcheck is installed"
else
    echo "   ℹ staticcheck not found (install with: go install honnef.co/go/tools/cmd/staticcheck@latest)"
fi

if command -v golangci-lint &> /dev/null; then
    echo "   ✓ golangci-lint is installed"
else
    echo "   ℹ golangci-lint not found (optional for linting)"
fi
echo

# Check Go module status
echo "6. Checking Go modules..."
cd "$(dirname "$0")"
if [ -f "go.mod" ]; then
    echo "   ✓ go.mod found"
    echo "   Running go mod verify..."
    if go mod verify &> /dev/null; then
        echo "   ✓ All modules verified successfully"
    else
        echo "   ✗ Module verification failed!"
    fi
else
    echo "   ✗ go.mod not found!"
fi
echo

# Check for fileb0x (asset embedding)
echo "7. Checking asset embedding tools..."
if go list -m github.com/Velocidex/fileb0x &> /dev/null 2>&1; then
    echo "   ✓ fileb0x module available"
else
    echo "   ✗ fileb0x module not found"
fi
echo

echo "=== Dependency Check Complete ==="
echo
echo "To install missing dependencies on Ubuntu/Debian:"
echo "  sudo apt-get install build-essential git gcc"
echo "  sudo apt-get install gcc-mingw-w64  # For Windows cross-compilation"
echo "  sudo apt-get install musl-tools     # For static Linux builds"
echo
echo "To install Go tools:"
echo "  go install honnef.co/go/tools/cmd/staticcheck@latest"
echo "  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"