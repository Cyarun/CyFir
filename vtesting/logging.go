package vtesting

import (
	"regexp"

	"github.com/Cyarun/CyFir/logging"
	"github.com/Cyarun/CyFir/vtesting/assert"
)

func MemoryLogsContain(t assert.TestingT, regex string, msgAndArgs ...interface{}) {
	if !MemoryLogsContainRegex(regex) {
		t.Errorf("Unable to find '%v' in memory logs %v", regex, msgAndArgs)
	}
}

func MemoryLogsContainRegex(regex string) bool {
	re := regexp.MustCompile(regex)

	for _, line := range logging.GetMemoryLogs() {
		if re.MatchString(line) {
			return true
		}
	}
	return false
}
