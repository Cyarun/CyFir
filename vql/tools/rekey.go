package tools

import (
	"context"
	"time"

	"github.com/Velocidex/ordereddict"
	"github.com/Cyarun/CyFir/acls"
	"github.com/Cyarun/CyFir/artifacts"
	config_proto "github.com/Cyarun/CyFir/config/proto"
	crypto_utils "github.com/Cyarun/CyFir/crypto/utils"
	"github.com/Cyarun/CyFir/services/writeback"
	"github.com/Cyarun/CyFir/vql"
	vql_subsystem "github.com/Cyarun/CyFir/vql"
	"www.velocidex.com/golang/vfilter"
	"www.velocidex.com/golang/vfilter/arg_parser"
)

var (
	ClientRestart = make(chan string)
)

type RekeyFunctionArgs struct {
	Wait int64 `vfilter:"optional,field=wait,doc=Wait this long before rekeying the client."`
}

type RekeyFunction struct{}

func (self *RekeyFunction) Call(ctx context.Context,
	scope vfilter.Scope,
	args *ordereddict.Dict) vfilter.Any {

	defer vql_subsystem.RegisterMonitor(ctx, "rekey", args)()

	arg := &RekeyFunctionArgs{}
	err := arg_parser.ExtractArgsWithContext(ctx, scope, args, arg)
	if err != nil {
		scope.Log("rekey: %v", err)
		return vfilter.Null{}
	}

	// This is a privileged operation
	err = vql_subsystem.CheckAccess(scope, acls.EXECVE)
	if err != nil {
		scope.Log("rekey: %v", err)
		return vfilter.Null{}
	}

	// Check the config if we are allowed to execve at all.
	client_config_obj, ok := artifacts.GetConfig(scope)
	if !ok || client_config_obj == nil {
		scope.Log("rekey: Must be running on a client to rekey")
		return vfilter.Null{}
	}

	config_obj := &config_proto.Config{Client: client_config_obj}
	pem, err := crypto_utils.GeneratePrivateKey()
	if err != nil {
		scope.Log("rekey: %v", err)
		return vfilter.Null{}
	}

	private_key, err := crypto_utils.ParseRsaPrivateKeyFromPemStr(pem)
	if err != nil {
		scope.Log("rekey: %v", err)
		return vfilter.Null{}
	}

	new_client_id := crypto_utils.ClientIDFromPublicKey(&private_key.PublicKey)

	// Update the write back.
	err = writeback.GetWritebackService().MutateWriteback(config_obj,
		func(wb *config_proto.Writeback) error {
			wb.PrivateKey = string(pem)
			wb.ClientId = new_client_id
			return writeback.WritebackUpdateLevel1
		})
	if err != nil {
		scope.Log("rekey: %v", err)
		return vfilter.Null{}
	}

	// Send the new client id to the main client loop so it can
	// restart, but wait a bit to allow messages to be sent to the
	// server on the old client id.
	go func() {
		time.Sleep(time.Duration(arg.Wait) * time.Second)

		select {
		case ClientRestart <- new_client_id:
		default:
		}
	}()
	return new_client_id
}

func (self RekeyFunction) Info(scope vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name:     "rekey",
		Doc:      "Causes the client to rekey and regenerate a new client ID. DANGEROUS! This will change the client's identity and it will appear as a new client in the GUI.",
		ArgType:  type_map.AddType(scope, &RekeyFunctionArgs{}),
		Metadata: vql.VQLMetadata().Permissions(acls.EXECVE).Build(),
	}
}

func init() {
	vql_subsystem.RegisterFunction(&RekeyFunction{})
}
