package main

import (
	config_proto "github.com/Cyarun/CyFir/config/proto"
)

// GetServiceName returns the appropriate service name based on configuration
func GetServiceName(config_obj *config_proto.Config) string {
	if config_obj.Client != nil && config_obj.Client.WindowsInstaller != nil {
		// Check if we should use legacy name
		if config_obj.Client.WindowsInstaller.ServiceName == "Velociraptor" {
			return "Velociraptor"
		}
		return config_obj.Client.WindowsInstaller.ServiceName
	}
	// Default to CyFir for new installations
	return "CyFir"
}

// GetServiceDescription returns the service description
func GetServiceDescription(config_obj *config_proto.Config) string {
	if config_obj.Client != nil && config_obj.Client.WindowsInstaller != nil &&
		config_obj.Client.WindowsInstaller.ServiceDescription != "" {
		return config_obj.Client.WindowsInstaller.ServiceDescription
	}
	return "CyFir - Cyber Forensics & IR Platform"
}