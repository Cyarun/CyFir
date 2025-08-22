//go:build windows
// +build windows

package logging

import config_proto "github.com/Cyarun/CyFir/config/proto"

// Syslog is not supported on Windows.
func maybeAddRemoteSyslog(
	config_obj *config_proto.Config, manager *LogManager) error {
	return nil
}
