//go:build !windows
// +build !windows

package config

func detectWindowsLegacy() bool {
	// Not on Windows
	return false
}