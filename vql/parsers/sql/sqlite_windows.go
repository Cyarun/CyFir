//go:build windows && amd64 && cgo
// +build windows,amd64,cgo

package sql

import (
	_ "github.com/Cyarun/CyFir/vql/windows/filesystems"
)
