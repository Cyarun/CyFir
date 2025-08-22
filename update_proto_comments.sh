#!/bin/bash
# Carefully update proto file comments for CyFir

echo "=== Updating Proto File Comments ==="
echo "This script only updates comments in proto files, not field names or messages"
echo

# Update config.proto comments
echo "Updating config.proto comments..."
sed -i 's|// Information about Velociraptor. This is a pseudo config item|// Information about CyFir. This is a pseudo config item|g' config/proto/config.proto
sed -i 's|// second). NOTE: Velociraptor typically holds|// second). NOTE: CyFir typically holds|g' config/proto/config.proto
sed -i 's|"NOTE: The self signed certificate must be signed by the Velociraptor CA|"NOTE: The self signed certificate must be signed by the CyFir CA|g' config/proto/config.proto
sed -i 's|// NOTE: Velociraptor has 2 layers of encryption|// NOTE: CyFir has 2 layers of encryption|g' config/proto/config.proto
sed -i 's|// A mapping between OIDC claim roles and Velociraptor roles|// A mapping between OIDC claim roles and CyFir roles|g' config/proto/config.proto
sed -i 's|// Client.Crypto.root_certs or the Velociraptor built in CA|// Client.Crypto.root_certs or the CyFir built in CA|g' config/proto/config.proto
sed -i 's|// A proxy setting to use - Velociraptor needs to connect|// A proxy setting to use - CyFir needs to connect|g' config/proto/config.proto
sed -i 's|// must be signed by the Velociraptor rooot CA|// must be signed by the CyFir root CA|g' config/proto/config.proto
sed -i 's|// Disable unicode usernames. By default Velociraptor allows|// Disable unicode usernames. By default CyFir allows|g' config/proto/config.proto
sed -i 's|// of the hex digits is ignored by Velociraptor|// of the hex digits is ignored by CyFir|g' config/proto/config.proto
sed -i 's|//   - THUMBPRINT_ONLY: Velociraptor only accepts|//   - THUMBPRINT_ONLY: CyFir only accepts|g' config/proto/config.proto
sed -i 's|// servers. This is required when connecting to Velociraptor with|// servers. This is required when connecting to CyFir with|g' config/proto/config.proto

# NOTE: Not changing field names like "VelociraptorServer" in quotes as those are protocol values

echo
echo "Checking other proto files for comment updates..."
find . -name "*.proto" -type f | while read proto_file; do
    if grep -q "// .*Velociraptor" "$proto_file" 2>/dev/null; then
        echo "Updating comments in: $proto_file"
        sed -i 's|// \(.*\)Velociraptor|// \1CyFir|g' "$proto_file"
    fi
done

echo
echo "=== Proto Comment Updates Complete ==="
echo "Note: Field names and string values were NOT changed to maintain protocol compatibility"