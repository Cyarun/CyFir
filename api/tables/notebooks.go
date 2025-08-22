package tables

import (
	"context"

	api_proto "github.com/Cyarun/CyFir/api/proto"
	config_proto "github.com/Cyarun/CyFir/config/proto"
	"github.com/Cyarun/CyFir/services"
)

func getNotebookTable(
	ctx context.Context,
	config_obj *config_proto.Config,
	in *api_proto.GetTableRequest,
	principal string) (*api_proto.GetTableResponse, error) {

	notebook_manager, err := services.GetNotebookManager(config_obj)
	if err != nil {
		return nil, err
	}

	_, err = notebook_manager.GetSharedNotebooks(ctx, principal)
	if err != nil {
		return nil, err
	}

	return getTable(ctx, config_obj, in, principal)
}
