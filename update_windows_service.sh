#!/bin/bash
# Update Windows service configurations for CyFir

echo "=== Updating Windows Service Configuration ==="

# Update server_service_windows.go
echo "Updating server_service_windows.go..."
sed -i 's/"server_service", "Manipulate the Velociraptor service."/"server_service", "Manipulate the CyFir service."/g' bin/server_service_windows.go
sed -i 's/"install", "Install Velociraptor frontend as a Windows service."/"install", "Install CyFir frontend as a Windows service."/g' bin/server_service_windows.go
sed -i 's/"remove", "Remove the Velociraptor Windows service."/"remove", "Remove the CyFir Windows service."/g' bin/server_service_windows.go
sed -i 's/name := "Velociraptor"/name := "CyFir"/g' bin/server_service_windows.go

# Note: Not changing function names like VelociraptorServerService to maintain API compatibility

# Update service installer for Darwin
echo "Updating Darwin installer..."
sed -i 's/Velociraptor service/CyFir service/g' bin/installer_darwin.go

# Update other service references
echo "Updating additional service files..."
find ./bin -name "*.go" -type f -exec grep -l '"Velociraptor"' {} \; | while read file; do
    # Only update user-visible strings in service context
    sed -i 's/"Installing Velociraptor service/"Installing CyFir service/g' "$file"
    sed -i 's/"Removing Velociraptor service/"Removing CyFir service/g' "$file"
    sed -i 's/"Starting Velociraptor service/"Starting CyFir service/g' "$file"
    sed -i 's/"Stopping Velociraptor service/"Stopping CyFir service/g' "$file"
done

echo "=== Windows Service Update Complete ==="