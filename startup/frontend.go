package startup

import (
	"context"

	"github.com/Cyarun/CyFir/api"
	config_proto "github.com/Cyarun/CyFir/config/proto"
	"github.com/Cyarun/CyFir/executor/throttler"
	"github.com/Cyarun/CyFir/logging"
	"github.com/Cyarun/CyFir/services"
	"github.com/Cyarun/CyFir/services/orgs"
	"github.com/Cyarun/CyFir/utils/tempfile"
	vql_subsystem "github.com/Cyarun/CyFir/vql"
	"github.com/Cyarun/CyFir/vql/networking"
)

// StartFrontendServices starts the binary as a frontend
func StartFrontendServices(
	ctx context.Context,
	config_obj *config_proto.Config) (*services.Service, error) {

	// Set the temp directory if needed
	tempfile.SetTempfile(config_obj)

	sm := services.NewServiceManager(ctx, config_obj)

	// Potentially restrict server functionality.
	err := MaybeEnforceAllowLists(config_obj)
	if err != nil {
		return sm, err
	}

	// Start throttling service
	err = sm.Start(throttler.StartStatsCollectorService)
	if err != nil {
		return sm, err
	}

	scope := vql_subsystem.MakeScope()
	scope.SetLogger(logging.NewPlainLogger(config_obj, &logging.FrontendComponent))

	vql_subsystem.InstallUnimplemented(scope)

	_, err = orgs.NewOrgManager(sm.Ctx, sm.Wg, config_obj)
	if err != nil {
		return sm, err
	}

	// Start the listening server
	server_builder, err := api.NewServerBuilder(sm.Ctx, config_obj, sm.Wg)
	if err != nil {
		return sm, err
	}

	err = networking.MaybeInstallDNSCache(sm.Ctx, sm.Wg, sm.Config)
	if err != nil {
		return sm, err
	}

	// Start the gRPC API server on the master only.
	if services.IsMaster(config_obj) {
		err = server_builder.WithAPIServer(sm.Ctx, sm.Wg)
		if err != nil {
			return sm, err
		}
	}

	return sm, server_builder.StartServer(sm.Ctx, sm.Wg)
}
