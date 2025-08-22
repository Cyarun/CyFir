package functions

import (
	"os"
	"testing"

	"github.com/Cyarun/CyFir/utils"
	"github.com/Cyarun/CyFir/vtesting/assert"
)

func TestEnvExpansion(t *testing.T) {
	os.Setenv("FOO_BAR", "Hello World")

	assert.Equal(t, "Hi, Hello World", utils.ExpandEnv("Hi, $FOO_BAR"))

	// Windows style expansion
	assert.Equal(t, "Hi, Hello World", utils.ExpandEnv("Hi, %FOO_BAR%"))

	// Can escape the $ char by doubling it
	assert.Equal(t, "Hi, $FOO_BAR", utils.ExpandEnv("Hi, $$FOO_BAR"))
}
