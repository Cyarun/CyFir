package data

import (
	"io/ioutil"
	"testing"

	"github.com/Velocidex/ordereddict"
	"github.com/Cyarun/CyFir/accessors"
	vql_subsystem "github.com/Cyarun/CyFir/vql"
	"github.com/Cyarun/CyFir/vtesting/assert"
)

func TestAccessorData(t *testing.T) {
	scope := vql_subsystem.MakeScope()
	accessor, err := accessors.GetAccessor("data", scope)
	assert.NoError(t, err)

	fd, err := accessor.Open("Hello world")
	assert.NoError(t, err)

	data, err := ioutil.ReadAll(fd)
	assert.NoError(t, err)

	assert.Equal(t, "Hello world", string(data))
}

func TestAccessorScope(t *testing.T) {
	scope := vql_subsystem.MakeScope().AppendVars(ordereddict.NewDict().
		Set("Foobar", "Hello world"))

	accessor, err := accessors.GetAccessor("scope", scope)
	assert.NoError(t, err)

	fd, err := accessor.Open("Foobar")
	assert.NoError(t, err)

	data, err := ioutil.ReadAll(fd)
	assert.NoError(t, err)

	assert.Equal(t, "Hello world", string(data))
}
