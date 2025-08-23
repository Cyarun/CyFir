//go:build windows
// +build windows

package config

import (
	"os"
	"golang.org/x/sys/windows/registry"
)

func detectWindowsLegacy() bool {
	// Check for Velociraptor service in registry
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, 
		`SYSTEM\CurrentControlSet\Services\Velociraptor`, registry.QUERY_VALUE)
	if err == nil {
		defer k.Close()
		return true
	}
	
	// Check for Velociraptor registry writeback location
	k, err = registry.OpenKey(registry.LOCAL_MACHINE,
		`SOFTWARE\Velocidex\Velociraptor`, registry.QUERY_VALUE)
	if err == nil {
		defer k.Close()
		return true
	}
	
	// Check for common installation paths
	legacyPaths := []string{
		`C:\Program Files\Velociraptor\Velociraptor.exe`,
		`C:\Program Files (x86)\Velociraptor\Velociraptor.exe`,
	}
	
	for _, path := range legacyPaths {
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}
	
	return false
}