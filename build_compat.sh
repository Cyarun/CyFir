#!/bin/bash
# Build compatibility script - creates both velociraptor and cyfir binaries

echo "Building with compatibility layer..."

# Build the binary as usual
make auto

# Create a copy with the new name
if [ -f "output/velociraptor" ]; then
    cp output/velociraptor output/cyfir
    echo "Created cyfir binary alongside velociraptor"
fi

if [ -f "output/velociraptor.exe" ]; then
    cp output/velociraptor.exe output/cyfir.exe
    echo "Created cyfir.exe binary alongside velociraptor.exe"  
fi