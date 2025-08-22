package functions

import (
	"context"

	"github.com/Velocidex/ordereddict"
	"github.com/Cyarun/CyFir/acls"
	"github.com/Cyarun/CyFir/utils"
	"github.com/Cyarun/CyFir/vql"
	vql_subsystem "github.com/Cyarun/CyFir/vql"
	"www.velocidex.com/golang/vfilter"
	"www.velocidex.com/golang/vfilter/arg_parser"
)

type ExpandPathArgs struct {
	Path string `vfilter:"required,field=path,doc=A path with environment escapes"`
}

type ExpandPath struct{}

func (self ExpandPath) Call(
	ctx context.Context,
	scope vfilter.Scope,
	args *ordereddict.Dict) vfilter.Any {

	defer vql_subsystem.RegisterMonitor(ctx, "expand", args)()

	err := vql_subsystem.CheckAccess(scope, acls.MACHINE_STATE)
	if err != nil {
		scope.Log("expand: %s", err)
		return vfilter.Null{}
	}

	arg := &ExpandPathArgs{}
	err = arg_parser.ExtractArgsWithContext(ctx, scope, args, arg)
	if err != nil {
		scope.Log("expand: %s", err.Error())
		return vfilter.Null{}
	}

	// Support windows style expansion on all platforms.
	return utils.ExpandEnv(arg.Path)
}

func (self ExpandPath) Info(scope vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name:     "expand",
		Doc:      "Expand the path using the environment.",
		ArgType:  type_map.AddType(scope, &ExpandPathArgs{}),
		Metadata: vql.VQLMetadata().Permissions(acls.MACHINE_STATE).Build(),
	}
}

func init() {
	vql_subsystem.RegisterFunction(&ExpandPath{})
}
