package config

import (
	config_proto "github.com/Cyarun/CyFir/config/proto"
)

// BrandingConfig allows gradual migration from Velociraptor to CyFir
type BrandingConfig struct {
	// Display name for UI
	DisplayName string
	
	// Full product name
	FullName string
	
	// Company name
	Company string
	
	// Service names
	WindowsServiceName string
	LinuxServiceName   string
	MacServiceName     string
	
	// Binary names
	ClientBinaryName string
	ServerBinaryName string
	
	// Package names
	RPMPackageName string
	DebPackageName string
	
	// Legacy mode - if true, use Velociraptor names
	LegacyMode bool
}

// GetBranding returns appropriate branding based on configuration
func GetBranding(config_obj *config_proto.Config) *BrandingConfig {
	// TODO: Enable when protobuf is regenerated
	// if config_obj != nil && config_obj.UseLegacyBranding {
	if false {
		return &BrandingConfig{
			DisplayName:        "Velociraptor",
			FullName:          "Velociraptor - Dig Deeper",
			Company:           "Rapid7 Inc",
			WindowsServiceName: "Velociraptor",
			LinuxServiceName:   "velociraptor",
			MacServiceName:     "com.velocidex.velociraptor",
			ClientBinaryName:   "velociraptor",
			ServerBinaryName:   "velociraptor",
			RPMPackageName:     "velociraptor",
			DebPackageName:     "velociraptor",
			LegacyMode:        true,
		}
	}
	
	return &BrandingConfig{
		DisplayName:        "CyFir",
		FullName:          "CyFir - Cyber Forensics & IR Platform",
		Company:           "CynorSense Solutions Pvt. Ltd.",
		WindowsServiceName: "CyFir",
		LinuxServiceName:   "cyfir",
		MacServiceName:     "com.cynorsense.cyfir",
		ClientBinaryName:   "cyfir",
		ServerBinaryName:   "cyfir",
		RPMPackageName:     "cyfir",
		DebPackageName:     "cyfir",
		LegacyMode:        false,
	}
}