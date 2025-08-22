#!/bin/bash
# Safe updates to vql.yaml documentation

echo "=== Updating VQL Documentation ==="

# Backup first
cp docs/references/vql.yaml docs/references/vql.yaml.bak

# Update user-visible documentation strings only
sed -i 's/Velociraptor comes with/CyFir comes with/g' docs/references/vql.yaml
sed -i 's/the Velociraptor is able/CyFir is able/g' docs/references/vql.yaml
sed -i "s/Internally, Velociraptor uses/Internally, CyFir uses/g" docs/references/vql.yaml
sed -i "s/Velociraptor maintains/CyFir maintains/g" docs/references/vql.yaml
sed -i "s/Velociraptor's external/CyFir's external/g" docs/references/vql.yaml
sed -i "s/Velociraptor does not/CyFir does not/g" docs/references/vql.yaml
sed -i "s/and Velociraptor's/and CyFir's/g" docs/references/vql.yaml

# Update example command lines
sed -i 's|"C:\\Program Files\\Velociraptor\\Velociraptor.exe"|"C:\\Program Files\\CyFir\\cyfir.exe"|g' docs/references/vql.yaml

# Keep VQL language name and technical terms unchanged
echo "Note: Keeping 'VQL' and other technical terms unchanged"

echo "=== VQL Documentation Update Complete ==="