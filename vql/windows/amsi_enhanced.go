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
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Velocidex/amsi"
	"github.com/Velocidex/ordereddict"
	"github.com/Cyarun/CyFir/acls"
	"github.com/Cyarun/CyFir/utils"
	"github.com/Cyarun/CyFir/vql"
	vql_subsystem "github.com/Cyarun/CyFir/vql"
	vfilter "www.velocidex.com/golang/vfilter"
	"www.velocidex.com/golang/vfilter/arg_parser"
)

const (
	AMSI_BULK_KEY = "$AMSI_BULK"
)

// AMSIScanResult represents the result of an AMSI scan
type AMSIScanResult struct {
	Input      string    `json:"input"`
	InputType  string    `json:"input_type"`
	Result     string    `json:"result"`
	IsMalware  bool      `json:"is_malware"`
	MD5Hash    string    `json:"md5_hash,omitempty"`
	Size       int64     `json:"size,omitempty"`
	ScanTime   time.Time `json:"scan_time"`
	ErrorMsg   string    `json:"error,omitempty"`
}

type _AMSIBulkScanArgs struct {
	Files       []string `vfilter:"optional,field=files,doc=List of files to scan"`
	Strings     []string `vfilter:"optional,field=strings,doc=List of strings to scan"`
	Directory   string   `vfilter:"optional,field=directory,doc=Directory to scan recursively"`
	Extensions  []string `vfilter:"optional,field=extensions,doc=File extensions to include (default: common script extensions)"`
	MaxFileSize int64    `vfilter:"optional,field=max_file_size,doc=Maximum file size to scan in bytes (default: 10MB)"`
	Timeout     int      `vfilter:"optional,field=timeout,doc=Timeout per file in seconds (default: 30)"`
}

type _AMSIBulkScanPlugin struct{}

func (self _AMSIBulkScanPlugin) Call(
	ctx context.Context,
	scope vfilter.Scope,
	args *ordereddict.Dict) <-chan vfilter.Row {
	output_chan := make(chan vfilter.Row)

	go func() {
		defer close(output_chan)
		defer vql_subsystem.RegisterMonitor(ctx, "amsi_bulk_scan", args)()

		err := vql_subsystem.CheckAccess(scope, acls.FILESYSTEM_READ)
		if err != nil {
			scope.Log("amsi_bulk_scan: %v", err)
			return
		}

		arg := &_AMSIBulkScanArgs{}
		err = arg_parser.ExtractArgsWithContext(ctx, scope, args, arg)
		if err != nil {
			scope.Log("amsi_bulk_scan: %v", err)
			return
		}

		// Set defaults
		if arg.MaxFileSize == 0 {
			arg.MaxFileSize = 10 * 1024 * 1024 // 10MB
		}
		if arg.Timeout == 0 {
			arg.Timeout = 30
		}
		if len(arg.Extensions) == 0 {
			arg.Extensions = []string{".ps1", ".bat", ".cmd", ".vbs", ".js", ".jse", ".wsf", ".wsh", ".scr", ".exe", ".dll"}
		}

		// Cache the AMSI session across the query context
		session_any := vql_subsystem.CacheGet(scope, AMSI_BULK_KEY)
		if session_any == nil {
			err := amsi.Initialize()
			if err != nil {
				scope.Log("amsi_bulk_scan: %v", err)
				return
			}
			session := amsi.OpenSession()

			// Tear it all down when the scope is destroyed
			vql_subsystem.GetRootScope(scope).AddDestructor(func() {
				amsi.CloseSession(session)
				amsi.Uninitialize()
			})
			vql_subsystem.CacheSet(scope, AMSI_BULK_KEY, session)
			session_any = session
		}

		session, ok := session_any.(*amsi.Session)
		if !ok {
			scope.Log("amsi_bulk_scan: invalid session")
			return
		}

		// Scan strings if provided
		for _, str := range arg.Strings {
			select {
			case <-ctx.Done():
				return
			default:
				result := scanString(session, str, scope)
				output_chan <- result
			}
		}

		// Scan individual files if provided
		for _, filePath := range arg.Files {
			select {
			case <-ctx.Done():
				return
			default:
				result := scanFile(session, filePath, arg.MaxFileSize, arg.Timeout, scope)
				output_chan <- result
			}
		}

		// Scan directory if provided
		if arg.Directory != "" {
			err := filepath.Walk(arg.Directory, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil // Continue on errors
				}

				// Skip directories
				if info.IsDir() {
					return nil
				}

				// Check file size
				if info.Size() > arg.MaxFileSize {
					return nil
				}

				// Check extension
				ext := strings.ToLower(filepath.Ext(path))
				includeFile := false
				for _, allowedExt := range arg.Extensions {
					if ext == strings.ToLower(allowedExt) {
						includeFile = true
						break
					}
				}

				if !includeFile {
					return nil
				}

				select {
				case <-ctx.Done():
					return fmt.Errorf("context cancelled")
				default:
					result := scanFile(session, path, arg.MaxFileSize, arg.Timeout, scope)
					select {
					case output_chan <- result:
					case <-ctx.Done():
						return fmt.Errorf("context cancelled")
					}
				}

				return nil
			})

			if err != nil {
				scope.Log("amsi_bulk_scan: directory walk error: %v", err)
			}
		}
	}()

	return output_chan
}

func (self _AMSIBulkScanPlugin) Info(scope vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.PluginInfo {
	return &vfilter.PluginInfo{
		Name:     "amsi_bulk_scan",
		Doc:      "Perform bulk AMSI scanning of files and strings for malware detection.",
		ArgType:  type_map.AddType(scope, &_AMSIBulkScanArgs{}),
		Metadata: vql.VQLMetadata().Permissions(acls.FILESYSTEM_READ).Build(),
	}
}

// scanString scans a string using AMSI
func scanString(session *amsi.Session, input string, scope vfilter.Scope) *AMSIScanResult {
	startTime := time.Now()
	
	result := &AMSIScanResult{
		Input:     input,
		InputType: "string",
		ScanTime:  startTime,
		Size:      int64(len(input)),
	}

	// Calculate MD5 hash
	hash := md5.Sum([]byte(input))
	result.MD5Hash = fmt.Sprintf("%x", hash)

	// Perform AMSI scan
	scanResult := session.ScanString(input)
	result.Result = convertAMSIResult(scanResult)
	result.IsMalware = (scanResult == amsi.ResultDetected)

	return result
}

// scanFile scans a file using AMSI
func scanFile(session *amsi.Session, filePath string, maxSize int64, timeout int, scope vfilter.Scope) *AMSIScanResult {
	startTime := time.Now()
	
	result := &AMSIScanResult{
		Input:     filePath,
		InputType: "file",
		ScanTime:  startTime,
	}

	// Check if file exists and get info
	info, err := os.Stat(filePath)
	if err != nil {
		result.ErrorMsg = fmt.Sprintf("File stat error: %v", err)
		return result
	}

	result.Size = info.Size()

	// Check file size
	if info.Size() > maxSize {
		result.ErrorMsg = fmt.Sprintf("File too large: %d bytes", info.Size())
		return result
	}

	// Read file content with timeout
	done := make(chan bool, 1)
	var content []byte
	var readErr error

	go func() {
		defer func() { done <- true }()
		
		file, err := os.Open(filePath)
		if err != nil {
			readErr = err
			return
		}
		defer file.Close()

		// Calculate MD5 while reading
		hash := md5.New()
		teeReader := io.TeeReader(file, hash)
		
		content, readErr = io.ReadAll(teeReader)
		if readErr == nil {
			result.MD5Hash = fmt.Sprintf("%x", hash.Sum(nil))
		}
	}()

	// Wait for read completion or timeout
	select {
	case <-done:
		if readErr != nil {
			result.ErrorMsg = fmt.Sprintf("File read error: %v", readErr)
			return result
		}
	case <-time.After(time.Duration(timeout) * time.Second):
		result.ErrorMsg = "File read timeout"
		return result
	}

	// Perform AMSI scan on content
	scanResult := session.ScanBuffer(content, filePath)
	result.Result = convertAMSIResult(scanResult)
	result.IsMalware = (scanResult == amsi.ResultDetected)

	return result
}

// convertAMSIResult converts AMSI result code to string
func convertAMSIResult(result amsi.AmsiResult) string {
	switch result {
	case amsi.ResultClean:
		return "Clean"
	case amsi.ResultNotDetected:
		return "NotDetected"
	case amsi.ResultBlockedByAdminStart:
		return "BlockedByAdminStart"
	case amsi.ResultBlockedByAdminEnd:
		return "BlockedByAdminEnd"
	case amsi.ResultDetected:
		return "Detected"
	default:
		return "Unknown"
	}
}

// AMSIMemoryScanArgs for memory scanning
type _AMSIMemoryScanArgs struct {
	ProcessID int64  `vfilter:"required,field=pid,doc=Process ID to scan"`
	Regions   []string `vfilter:"optional,field=regions,doc=Memory regions to scan (default: all executable regions)"`
}

type _AMSIMemoryScanPlugin struct{}

func (self _AMSIMemoryScanPlugin) Call(
	ctx context.Context,
	scope vfilter.Scope,
	args *ordereddict.Dict) <-chan vfilter.Row {
	output_chan := make(chan vfilter.Row)

	go func() {
		defer close(output_chan)
		defer vql_subsystem.RegisterMonitor(ctx, "amsi_memory_scan", args)()

		err := vql_subsystem.CheckAccess(scope, acls.MACHINE_STATE)
		if err != nil {
			scope.Log("amsi_memory_scan: %v", err)
			return
		}

		arg := &_AMSIMemoryScanArgs{}
		err = arg_parser.ExtractArgsWithContext(ctx, scope, args, arg)
		if err != nil {
			scope.Log("amsi_memory_scan: %v", err)
			return
		}

		// Cache the AMSI session across the query context
		session_any := vql_subsystem.CacheGet(scope, AMSI_BULK_KEY)
		if session_any == nil {
			err := amsi.Initialize()
			if err != nil {
				scope.Log("amsi_memory_scan: %v", err)
				return
			}
			session := amsi.OpenSession()

			vql_subsystem.GetRootScope(scope).AddDestructor(func() {
				amsi.CloseSession(session)
				amsi.Uninitialize()
			})
			vql_subsystem.CacheSet(scope, AMSI_BULK_KEY, session)
			session_any = session
		}

		session, ok := session_any.(*amsi.Session)
		if !ok {
			scope.Log("amsi_memory_scan: invalid session")
			return
		}

		// Scan process memory
		results := scanProcessMemory(session, arg.ProcessID, scope)
		for _, result := range results {
			select {
			case <-ctx.Done():
				return
			case output_chan <- result:
			}
		}
	}()

	return output_chan
}

func (self _AMSIMemoryScanPlugin) Info(scope vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.PluginInfo {
	return &vfilter.PluginInfo{
		Name:     "amsi_memory_scan",
		Doc:      "Scan process memory using AMSI for in-memory threats.",
		ArgType:  type_map.AddType(scope, &_AMSIMemoryScanArgs{}),
		Metadata: vql.VQLMetadata().Permissions(acls.MACHINE_STATE).Build(),
	}
}

// scanProcessMemory scans process memory regions
func scanProcessMemory(session *amsi.Session, pid int64, scope vfilter.Scope) []*AMSIScanResult {
	var results []*AMSIScanResult

	// This is a simplified implementation - in practice, you would use
	// process memory APIs to read memory regions and scan them
	scope.Log("amsi_memory_scan: Memory scanning for PID %d not yet fully implemented", pid)
	
	// Placeholder result
	result := &AMSIScanResult{
		Input:     fmt.Sprintf("Process %d memory", pid),
		InputType: "memory",
		ScanTime:  time.Now(),
		Result:    "NotImplemented",
		ErrorMsg:  "Memory scanning feature is not yet fully implemented",
	}
	
	results = append(results, result)
	return results
}

func init() {
	vql_subsystem.RegisterPlugin(&_AMSIBulkScanPlugin{})
	vql_subsystem.RegisterPlugin(&_AMSIMemoryScanPlugin{})
}