package paths

import (
	"strings"

	"github.com/Cyarun/CyFir/file_store/api"
)

// Convert an artifact name to a path component to store its definition.
func GetArtifactDefintionPath(name string) api.FSPathSpec {
	return ARTIFACT_DEFINITION_PREFIX.
		AddUnsafeChild(strings.Split(name, ".")...)
}
