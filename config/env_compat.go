package config

import (
	"os"
)

// GetEnvWithFallback checks for both new (CYFIR_*) and old (VELOCIRAPTOR_*) 
// environment variables to maintain backward compatibility
func GetEnvWithFallback(newName, oldName string) string {
	// First check the new CYFIR_* variable
	if value := os.Getenv(newName); value != "" {
		return value
	}
	// Fall back to the old VELOCIRAPTOR_* variable
	return os.Getenv(oldName)
}

// Environment variable compatibility mapping
func GetConfigEnv() string {
	return GetEnvWithFallback("CYFIR_CONFIG", "VELOCIRAPTOR_CONFIG")
}

func GetLiteralConfigEnv() string {
	return GetEnvWithFallback("CYFIR_LITERAL_CONFIG", "VELOCIRAPTOR_LITERAL_CONFIG")
}

func GetAPIConfigEnv() string {
	return GetEnvWithFallback("CYFIR_API_CONFIG", "VELOCIRAPTOR_API_CONFIG")
}