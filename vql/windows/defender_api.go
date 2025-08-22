//go:build windows
// +build windows

/*
   Velociraptor - Dig Deeper
   Copyright (C) 2019-2025 Rapid7 Inc.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package windows

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Velocidex/ordereddict"
	"golang.org/x/sys/windows/registry"
	"github.com/Cyarun/CyFir/acls"
	"github.com/Cyarun/CyFir/vql"
	vql_subsystem "github.com/Cyarun/CyFir/vql"
	vfilter "www.velocidex.com/golang/vfilter"
	"www.velocidex.com/golang/vfilter/arg_parser"
)

// DefenderStatus represents Windows Defender status information
type DefenderStatus struct {
	RealTimeProtectionEnabled bool      `json:"real_time_protection_enabled"`
	BehaviorMonitorEnabled    bool      `json:"behavior_monitor_enabled"`
	OnAccessProtectionEnabled bool      `json:"on_access_protection_enabled"`
	IoavProtectionEnabled     bool      `json:"ioav_protection_enabled"`
	TamperProtectionEnabled   bool      `json:"tamper_protection_enabled"`
	CloudProtectionLevel      int       `json:"cloud_protection_level"`
	SubmitSamplesConsent      int       `json:"submit_samples_consent"`
	AntivirusSignatureVersion string    `json:"antivirus_signature_version"`
	AntispywareSignatureAge   int       `json:"antispyware_signature_age"`
	LastFullScanTime          time.Time `json:"last_full_scan_time"`
	LastQuickScanTime         time.Time `json:"last_quick_scan_time"`
	ServiceRunning            bool      `json:"service_running"`
	ServiceStartType          string    `json:"service_start_type"`
}

// DefenderExclusion represents exclusion settings
type DefenderExclusion struct {
	Type  string   `json:"type"`
	Items []string `json:"items"`
}

// DefenderThreatInfo represents threat information
type DefenderThreatInfo struct {
	ThreatID         string    `json:"threat_id"`
	ThreatName       string    `json:"threat_name"`
	SeverityID       int       `json:"severity_id"`
	CategoryID       int       `json:"category_id"`
	Path             string    `json:"path"`
	DetectionTime    time.Time `json:"detection_time"`
	ActionTaken      string    `json:"action_taken"`
	IsActiveDetection bool     `json:"is_active_detection"`
}

type DefenderStatusArgs struct{}

type DefenderStatusPlugin struct{}

func (self DefenderStatusPlugin) Call(
	ctx context.Context,
	scope vfilter.Scope,
	args *ordereddict.Dict) <-chan vfilter.Row {
	output_chan := make(chan vfilter.Row)

	go func() {
		defer close(output_chan)
		defer vql_subsystem.RegisterMonitor(ctx, "defender_status", args)()

		err := vql_subsystem.CheckAccess(scope, acls.MACHINE_STATE)
		if err != nil {
			scope.Log("defender_status: %v", err)
			return
		}

		arg := &DefenderStatusArgs{}
		err = arg_parser.ExtractArgsWithContext(ctx, scope, args, arg)
		if err != nil {
			scope.Log("defender_status: %v", err)
			return
		}

		status := getDefenderStatus(scope)

		select {
		case <-ctx.Done():
			return
		case output_chan <- status:
		}
	}()

	return output_chan
}

func (self DefenderStatusPlugin) Info(scope vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.PluginInfo {
	return &vfilter.PluginInfo{
		Name:     "defender_status",
		Doc:      "Get Windows Defender status and configuration.",
		ArgType:  type_map.AddType(scope, &DefenderStatusArgs{}),
		Metadata: vql.VQLMetadata().Permissions(acls.MACHINE_STATE).Build(),
	}
}

// DefenderExclusionsArgs for the exclusions plugin
type DefenderExclusionsArgs struct{}

type DefenderExclusionsPlugin struct{}

func (self DefenderExclusionsPlugin) Call(
	ctx context.Context,
	scope vfilter.Scope,
	args *ordereddict.Dict) <-chan vfilter.Row {
	output_chan := make(chan vfilter.Row)

	go func() {
		defer close(output_chan)
		defer vql_subsystem.RegisterMonitor(ctx, "defender_exclusions", args)()

		err := vql_subsystem.CheckAccess(scope, acls.MACHINE_STATE)
		if err != nil {
			scope.Log("defender_exclusions: %v", err)
			return
		}

		arg := &DefenderExclusionsArgs{}
		err = arg_parser.ExtractArgsWithContext(ctx, scope, args, arg)
		if err != nil {
			scope.Log("defender_exclusions: %v", err)
			return
		}

		exclusions := getDefenderExclusions(scope)
		for _, exclusion := range exclusions {
			select {
			case <-ctx.Done():
				return
			case output_chan <- exclusion:
			}
		}
	}()

	return output_chan
}

func (self DefenderExclusionsPlugin) Info(scope vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.PluginInfo {
	return &vfilter.PluginInfo{
		Name:     "defender_exclusions",
		Doc:      "Get Windows Defender exclusion settings.",
		ArgType:  type_map.AddType(scope, &DefenderExclusionsArgs{}),
		Metadata: vql.VQLMetadata().Permissions(acls.MACHINE_STATE).Build(),
	}
}

// DefenderThreatsArgs for the threats plugin
type DefenderThreatsArgs struct {
	ActiveOnly bool `vfilter:"optional,field=active_only,doc=Only return active threats"`
}

type DefenderThreatsPlugin struct{}

func (self DefenderThreatsPlugin) Call(
	ctx context.Context,
	scope vfilter.Scope,
	args *ordereddict.Dict) <-chan vfilter.Row {
	output_chan := make(chan vfilter.Row)

	go func() {
		defer close(output_chan)
		defer vql_subsystem.RegisterMonitor(ctx, "defender_threats", args)()

		err := vql_subsystem.CheckAccess(scope, acls.MACHINE_STATE)
		if err != nil {
			scope.Log("defender_threats: %v", err)
			return
		}

		arg := &DefenderThreatsArgs{}
		err = arg_parser.ExtractArgsWithContext(ctx, scope, args, arg)
		if err != nil {
			scope.Log("defender_threats: %v", err)
			return
		}

		threats := getDefenderThreats(scope, arg.ActiveOnly)
		for _, threat := range threats {
			select {
			case <-ctx.Done():
				return
			case output_chan <- threat:
			}
		}
	}()

	return output_chan
}

func (self DefenderThreatsPlugin) Info(scope vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.PluginInfo {
	return &vfilter.PluginInfo{
		Name:     "defender_threats",
		Doc:      "Get Windows Defender threat detection information.",
		ArgType:  type_map.AddType(scope, &DefenderThreatsArgs{}),
		Metadata: vql.VQLMetadata().Permissions(acls.MACHINE_STATE).Build(),
	}
}

// Helper functions to read Windows Defender information

func getDefenderStatus(scope vfilter.Scope) *DefenderStatus {
	status := &DefenderStatus{}

	// Read Real-Time Protection settings
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, 
		`SOFTWARE\Microsoft\Windows Defender\Real-Time Protection`, registry.QUERY_VALUE)
	if err == nil {
		defer key.Close()

		if val, _, err := key.GetIntegerValue("DisableRealtimeMonitoring"); err == nil {
			status.RealTimeProtectionEnabled = val == 0
		}
		if val, _, err := key.GetIntegerValue("DisableBehaviorMonitoring"); err == nil {
			status.BehaviorMonitorEnabled = val == 0
		}
		if val, _, err := key.GetIntegerValue("DisableOnAccessProtection"); err == nil {
			status.OnAccessProtectionEnabled = val == 0
		}
		if val, _, err := key.GetIntegerValue("DisableIOAVProtection"); err == nil {
			status.IoavProtectionEnabled = val == 0
		}
	}

	// Read Features settings for Tamper Protection
	key, err = registry.OpenKey(registry.LOCAL_MACHINE, 
		`SOFTWARE\Microsoft\Windows Defender\Features`, registry.QUERY_VALUE)
	if err == nil {
		defer key.Close()

		if val, _, err := key.GetIntegerValue("TamperProtection"); err == nil {
			status.TamperProtectionEnabled = val == 1
		}
	}

	// Read SpyNet (Cloud Protection) settings
	key, err = registry.OpenKey(registry.LOCAL_MACHINE, 
		`SOFTWARE\Microsoft\Windows Defender\SpyNet`, registry.QUERY_VALUE)
	if err == nil {
		defer key.Close()

		if val, _, err := key.GetIntegerValue("SpyNetReporting"); err == nil {
			status.CloudProtectionLevel = int(val)
		}
		if val, _, err := key.GetIntegerValue("SubmitSamplesConsent"); err == nil {
			status.SubmitSamplesConsent = int(val)
		}
	}

	// Read signature information
	key, err = registry.OpenKey(registry.LOCAL_MACHINE, 
		`SOFTWARE\Microsoft\Windows Defender\Signature Updates`, registry.QUERY_VALUE)
	if err == nil {
		defer key.Close()

		if val, _, err := key.GetStringValue("AVSignatureVersion"); err == nil {
			status.AntivirusSignatureVersion = val
		}
		if val, _, err := key.GetIntegerValue("SignatureUpdateInterval"); err == nil {
			status.AntispywareSignatureAge = int(val)
		}
	}

	// Read scan times
	key, err = registry.OpenKey(registry.LOCAL_MACHINE, 
		`SOFTWARE\Microsoft\Windows Defender\Scan`, registry.QUERY_VALUE)
	if err == nil {
		defer key.Close()

		if val, _, err := key.GetIntegerValue("LastFullScanTime"); err == nil {
			status.LastFullScanTime = time.Unix(int64(val), 0)
		}
		if val, _, err := key.GetIntegerValue("LastQuickScanTime"); err == nil {
			status.LastQuickScanTime = time.Unix(int64(val), 0)
		}
	}

	// Check service status
	serviceKey, err := registry.OpenKey(registry.LOCAL_MACHINE, 
		`SYSTEM\CurrentControlSet\Services\WinDefend`, registry.QUERY_VALUE)
	if err == nil {
		defer serviceKey.Close()

		if val, _, err := serviceKey.GetIntegerValue("Start"); err == nil {
			switch val {
			case 0:
				status.ServiceStartType = "Boot"
			case 1:
				status.ServiceStartType = "System"
			case 2:
				status.ServiceStartType = "Automatic"
			case 3:
				status.ServiceStartType = "Manual"
			case 4:
				status.ServiceStartType = "Disabled"
			default:
				status.ServiceStartType = "Unknown"
			}
		}
	}

	return status
}

func getDefenderExclusions(scope vfilter.Scope) []*DefenderExclusion {
	var exclusions []*DefenderExclusion

	exclusionTypes := map[string]string{
		"Paths":      "Path",
		"Extensions": "Extension", 
		"Processes":  "Process",
	}

	for regName, typeName := range exclusionTypes {
		key, err := registry.OpenKey(registry.LOCAL_MACHINE, 
			fmt.Sprintf(`SOFTWARE\Microsoft\Windows Defender\Exclusions\%s`, regName), 
			registry.QUERY_VALUE)
		if err != nil {
			continue
		}
		defer key.Close()

		valueNames, err := key.ReadValueNames(-1)
		if err != nil {
			continue
		}

		if len(valueNames) > 0 {
			exclusions = append(exclusions, &DefenderExclusion{
				Type:  typeName,
				Items: valueNames,
			})
		}
	}

	return exclusions
}

func getDefenderThreats(scope vfilter.Scope, activeOnly bool) []*DefenderThreatInfo {
	var threats []*DefenderThreatInfo

	// Read from Detection History in registry
	detectionKey, err := registry.OpenKey(registry.LOCAL_MACHINE, 
		`SOFTWARE\Microsoft\Windows Defender\Threats`, registry.QUERY_VALUE)
	if err != nil {
		return threats
	}
	defer detectionKey.Close()

	subKeys, err := detectionKey.ReadSubKeyNames(-1)
	if err != nil {
		return threats
	}

	for _, threatID := range subKeys {
		threatKey, err := registry.OpenKey(detectionKey, threatID, registry.QUERY_VALUE)
		if err != nil {
			continue
		}

		threat := &DefenderThreatInfo{ThreatID: threatID}

		// Read threat information
		if val, _, err := threatKey.GetStringValue("ThreatName"); err == nil {
			threat.ThreatName = val
		}
		if val, _, err := threatKey.GetIntegerValue("SeverityID"); err == nil {
			threat.SeverityID = int(val)
		}
		if val, _, err := threatKey.GetIntegerValue("CategoryID"); err == nil {
			threat.CategoryID = int(val)
		}
		if val, _, err := threatKey.GetStringValue("Path"); err == nil {
			threat.Path = val
		}
		if val, _, err := threatKey.GetStringValue("DetectionTime"); err == nil {
			if detectionTime, err := strconv.ParseInt(val, 10, 64); err == nil {
				threat.DetectionTime = time.Unix(detectionTime, 0)
			}
		}
		if val, _, err := threatKey.GetStringValue("ActionTaken"); err == nil {
			threat.ActionTaken = val
		}
		if val, _, err := threatKey.GetIntegerValue("IsActive"); err == nil {
			threat.IsActiveDetection = val == 1
		}

		threatKey.Close()

		// Filter by active status if requested
		if !activeOnly || threat.IsActiveDetection {
			threats = append(threats, threat)
		}
	}

	return threats
}

func init() {
	vql_subsystem.RegisterPlugin(&DefenderStatusPlugin{})
	vql_subsystem.RegisterPlugin(&DefenderExclusionsPlugin{})
	vql_subsystem.RegisterPlugin(&DefenderThreatsPlugin{})
}