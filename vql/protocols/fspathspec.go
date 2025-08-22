package protocols

import (
	"github.com/Cyarun/CyFir/file_store/path_specs"
	vql_subsystem "github.com/Cyarun/CyFir/vql"
	"www.velocidex.com/golang/vfilter"
)

type _PathspecRegex struct{}

func (self _PathspecRegex) Applicable(a vfilter.Any, b vfilter.Any) bool {
	switch b.(type) {
	case path_specs.FSPathSpec, path_specs.DSPathSpec:
	default:
		return false
	}
	_, ok := a.(string)
	return ok
}

func (self _PathspecRegex) Match(scope vfilter.Scope, a vfilter.Any, b vfilter.Any) bool {
	lhs := ""
	switch t := b.(type) {
	case path_specs.FSPathSpec:
		lhs = t.AsClientPath()
	case path_specs.DSPathSpec:
		lhs = t.AsClientPath()
	default:
		return false
	}

	return scope.Match(a, lhs)
}

func init() {
	vql_subsystem.RegisterProtocol(&_PathspecRegex{})
}
