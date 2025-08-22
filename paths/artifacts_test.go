package paths_test

import (
	"github.com/Cyarun/CyFir/paths"
	"github.com/Cyarun/CyFir/vtesting/assert"
)

func (self *PathManagerTestSuite) TestArtifactPathManager() {
	assert.Equal(self.T(),
		"/fs/artifact_definitions/Windows/Some/Artifact.yaml",
		self.getFilestorePath(
			paths.GetArtifactDefintionPath("Windows.Some.Artifact")))

}
