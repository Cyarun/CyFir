package file

import (
	"github.com/Cyarun/CyFir/accessors"
	"github.com/Cyarun/CyFir/acls"
	"github.com/Cyarun/CyFir/utils"
)

var (
	AllowedPrefixes *utils.PrefixTree
	DeniedError     = utils.Wrap(acls.PermissionDenied, "No accesss to filesystem path")
)

func CheckPath(full_path string) error {
	destination_path, err := accessors.NewNativePath(full_path)
	if err != nil || len(destination_path.Components) == 0 {
		return err
	}

	return CheckPrefix(destination_path)
}

func CheckPrefix(full_path *accessors.OSPath) error {
	// All files are allowed
	if AllowedPrefixes == nil {
		return nil
	}

	if len(full_path.Components) == 0 {
		return nil
	}

	if AllowedPrefixes.Present(full_path.Components) {
		return nil
	}

	return DeniedError
}
