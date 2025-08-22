#!/bin/bash

# CyFir Service Installation Compatibility Script
# Supports both legacy (Velociraptor) and new (CyFir) service names

BINARY_PATH="${1:-/usr/local/bin/cyfir}"
CONFIG_PATH="${2:-/etc/cyfir/server.config.yaml}"
SERVICE_NAME="${3:-cyfir}"
LEGACY_MODE="${4:-false}"

echo "=== CyFir Service Installation ==="
echo "Binary: $BINARY_PATH"
echo "Config: $CONFIG_PATH"
echo "Service: $SERVICE_NAME"
echo "Legacy Mode: $LEGACY_MODE"
echo

# Detect OS
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # Linux with systemd
    if command -v systemctl &> /dev/null; then
        echo "Installing systemd service..."
        
        # Create service file
        SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"
        
        cat > "$SERVICE_FILE" << EOF
[Unit]
Description=CyFir - Cyber Forensics & IR Platform
After=network.target

[Service]
Type=simple
User=root
ExecStart=$BINARY_PATH --config $CONFIG_PATH frontend
Restart=always
RestartSec=10
KillMode=process

[Install]
WantedBy=multi-user.target
EOF
        
        # If legacy mode, also create velociraptor service as alias
        if [ "$LEGACY_MODE" = "true" ]; then
            ln -sf "$SERVICE_FILE" "/etc/systemd/system/velociraptor.service"
            echo "Created legacy service alias: velociraptor.service"
        fi
        
        # Enable and start service
        systemctl daemon-reload
        systemctl enable "$SERVICE_NAME"
        echo "✓ Service installed: $SERVICE_NAME"
        
    else
        echo "⚠ Systemd not found. Manual service setup required."
    fi
    
elif [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS with launchd
    echo "Installing launchd service..."
    
    PLIST_PATH="/Library/LaunchDaemons/com.cynorsense.${SERVICE_NAME}.plist"
    
    cat > "$PLIST_PATH" << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.cynorsense.${SERVICE_NAME}</string>
    <key>ProgramArguments</key>
    <array>
        <string>$BINARY_PATH</string>
        <string>--config</string>
        <string>$CONFIG_PATH</string>
        <string>frontend</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
    <key>UserName</key>
    <string>root</string>
</dict>
</plist>
EOF
    
    chown root:wheel "$PLIST_PATH"
    chmod 644 "$PLIST_PATH"
    
    # Load service
    launchctl load "$PLIST_PATH"
    echo "✓ Service installed: com.cynorsense.${SERVICE_NAME}"
    
else
    echo "⚠ Unsupported OS: $OSTYPE"
    exit 1
fi

echo
echo "=== Installation Complete ==="
echo "Service name: $SERVICE_NAME"
echo "Start with: systemctl start $SERVICE_NAME (Linux)"
echo "Check status: systemctl status $SERVICE_NAME (Linux)"