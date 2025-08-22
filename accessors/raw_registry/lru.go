package raw_registry

import (
	"www.velocidex.com/golang/regparser"
	"github.com/Cyarun/CyFir/accessors"
)

type readDirLRUItem struct {
	children []accessors.FileInfo
	err      error

	key *regparser.CM_KEY_NODE
}
