//go:build windows
// +build windows

package registry

import (
	"time"

	"github.com/Cyarun/CyFir/accessors"
)

type readDirLRUItem struct {
	children []accessors.FileInfo
	err      error
	age      time.Time
}
