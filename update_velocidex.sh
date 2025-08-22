#!/bin/bash

# Update Velocidex references to CynorSense

echo "Updating Velocidex references to CynorSense..."

# Update www.velocidex.com to cynorsense.com  
find . -name "*.go" -type f -exec sed -i 's/www\.velocidex\.com/cynorsense.com/g' {} \;
find . -name "*.yaml" -type f -exec sed -i 's/www\.velocidex\.com/cynorsense.com/g' {} \;
find . -name "*.yml" -type f -exec sed -i 's/www\.velocidex\.com/cynorsense.com/g' {} \;
find . -name "*.md" -type f -exec sed -i 's/www\.velocidex\.com/cynorsense.com/g' {} \;

# Update Velocidex Enterprises to CynorSense Solutions
find . -name "*.go" -type f -exec sed -i 's/Velocidex Enterprises/CynorSense Solutions/g' {} \;
find . -name "*.yaml" -type f -exec sed -i 's/Velocidex Enterprises/CynorSense Solutions/g' {} \;
find . -name "*.yml" -type f -exec sed -i 's/Velocidex Enterprises/CynorSense Solutions/g' {} \;
find . -name "*.md" -type f -exec sed -i 's/Velocidex Enterprises/CynorSense Solutions/g' {} \;

# Update support@velocidex.com to support@cynorsense.com
find . -name "*.go" -type f -exec sed -i 's/support@velocidex\.com/support@cynorsense.com/g' {} \;
find . -name "*.yaml" -type f -exec sed -i 's/support@velocidex\.com/support@cynorsense.com/g' {} \;
find . -name "*.yml" -type f -exec sed -i 's/support@velocidex\.com/support@cynorsense.com/g' {} \;
find . -name "*.md" -type f -exec sed -i 's/support@velocidex\.com/support@cynorsense.com/g' {} \;

echo "Velocidex update complete!"