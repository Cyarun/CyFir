package artifacts

import (
	config_proto "github.com/Cyarun/CyFir/config/proto"
	"github.com/Cyarun/CyFir/constants"
	"github.com/Cyarun/CyFir/uploads"
	"github.com/Cyarun/CyFir/utils"
	vql_subsystem "github.com/Cyarun/CyFir/vql"
	"www.velocidex.com/golang/vfilter"
)

// Gets the client config from the scope.
func GetConfig(scope vfilter.Scope) (*config_proto.ClientConfig, bool) {
	scope_config, pres := scope.Resolve(constants.SCOPE_CONFIG)
	if !pres {
		return nil, false
	}

	config, ok := scope_config.(*config_proto.ClientConfig)
	if config == nil {
		return nil, false
	}
	return config, ok
}

func GetUploader(scope vfilter.Scope) (uploads.Uploader, bool) {
	scope_uploader, pres := scope.Resolve(constants.SCOPE_UPLOADER)
	if !pres {
		return nil, false
	}

	config, ok := scope_uploader.(uploads.Uploader)
	if utils.IsNil(config) {
		return nil, false
	}

	return config, ok
}

func GetACLManager(scope vfilter.Scope) (vql_subsystem.ACLManager, bool) {
	scope_manager, pres := scope.Resolve(vql_subsystem.ACL_MANAGER_VAR)
	if !pres {
		return nil, false
	}

	config, ok := scope_manager.(vql_subsystem.ACLManager)
	if utils.IsNil(config) {
		return nil, false
	}

	return config, ok
}
