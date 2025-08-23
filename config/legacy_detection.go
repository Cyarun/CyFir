package config

import (
	"os"
	"runtime"
	
	config_proto "github.com/Cyarun/CyFir/config/proto"
)

// DetectLegacyInstallation checks for existing Velociraptor installations
// and returns true if legacy branding should be used
func DetectLegacyInstallation() bool {
	switch runtime.GOOS {
	case "windows":
		return detectWindowsLegacy()
	case "linux":
		return detectLinuxLegacy()
	case "darwin":
		return detectDarwinLegacy()
	default:
		return false
	}
}

// detectWindowsLegacy is implemented in legacy_detection_windows.go for Windows
// and legacy_detection_unix.go for other platforms

func detectLinuxLegacy() bool {
	// Check for systemd service
	servicePaths := []string{
		"/etc/systemd/system/velociraptor.service",
		"/lib/systemd/system/velociraptor.service",
	}
	
	for _, path := range servicePaths {
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}
	
	// Check for common installation paths
	legacyPaths := []string{
		"/usr/local/bin/velociraptor",
		"/usr/bin/velociraptor",
		"/opt/velociraptor/velociraptor",
	}
	
	for _, path := range legacyPaths {
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}
	
	// Check for writeback file
	if _, err := os.Stat("/etc/velociraptor.writeback.yaml"); err == nil {
		return true
	}
	
	return false
}

func detectDarwinLegacy() bool {
	// Check for launchd service
	servicePaths := []string{
		"/Library/LaunchDaemons/com.velocidex.velociraptor.plist",
		"/System/Library/LaunchDaemons/com.velocidex.velociraptor.plist",
	}
	
	for _, path := range servicePaths {
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}
	
	// Check for common installation paths
	legacyPaths := []string{
		"/usr/local/sbin/velociraptor",
		"/usr/local/bin/velociraptor",
	}
	
	for _, path := range legacyPaths {
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}
	
	return false
}

// ApplyBrandingDetection updates the config with detected legacy mode
func ApplyBrandingDetection(config_obj *config_proto.Config) {
	// TODO: Enable when protobuf is regenerated
	// if !config_obj.UseLegacyBranding {
	//     config_obj.UseLegacyBranding = DetectLegacyInstallation()
	// }
}