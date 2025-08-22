package main

import (
	"os"
	"strings"

	"github.com/Cyarun/CyFir/config"
	logging "github.com/Cyarun/CyFir/logging"
)

var (
	nobanner_flag = app.Flag(
		"nobanner", "Suppress the CyFir banner").Bool()
)

var banner = `
<green>  ______      _______ _      
<green> / ___\ \    / /  ___(_)     
<green>| |    \ \  / /| |_   _ _ __  
<green>| |     \ \/ / |  _| | | '__| 
<green>| |___   \  /  | |   | | |    
<green> \____|   \/   |_|   |_|_|    
<green>                              
<red>Cyber Forensics & IR Platform  <cyan>by CynorSense Solutions
`

func doBanner() {
	if *nobanner_flag {
		return
	}
	for _, line := range strings.Split(banner, "\n") {
		if len(line) > 0 {
			logging.Prelog(line)
		}
	}

	version := config.GetVersion()

	logging.Prelog("<yellow>This is CyFir %v built on %v (%v)", version.Version,
		version.BuildTime, version.Commit)

	// Record some important environment variables to the log.
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) != 2 {
			continue
		}

		switch pair[0] {
		case // Ignore this one as it is the actual configuration
			"VELOCIRAPTOR_LITERAL_CONFIG",
			"VELOCIRAPTOR_SLOW_FILESYSTEM",
			"VELOCIRAPTOR_DISABLE_CSRF", "VELOCIRAPTOR_INJECT_API_SLEEP",
			"GOGC", "GOTRACEBACK", "GOMAXPROCS", "GODEBUG", "GOMEMLIMIT",
			"GORACE":
			logging.Prelog("<yellow>Environment Variable %v:</> %v",
				pair[0], pair[1])
		}
	}
}
