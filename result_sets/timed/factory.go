package timed

import (
	"context"

	"github.com/Velocidex/json"
	config_proto "github.com/Cyarun/CyFir/config/proto"
	"github.com/Cyarun/CyFir/file_store/api"
	"github.com/Cyarun/CyFir/result_sets"

	_ "github.com/Cyarun/CyFir/result_sets/simple"
)

type TimedFactory struct{}

func (self TimedFactory) NewTimedResultSetWriter(
	config_obj *config_proto.Config,
	path_manager api.PathManager,
	opts *json.EncOpts,
	completion func()) (result_sets.TimedResultSetWriter, error) {
	return NewTimedResultSetWriter(
		config_obj, path_manager, opts, completion)
}

func (self TimedFactory) NewTimedResultSetReader(
	ctx context.Context,
	config_obj *config_proto.Config,
	path_manager api.PathManager) (result_sets.TimedResultSetReader, error) {

	return &TimedResultSetReader{
		files:      path_manager.GetAvailableFiles(ctx),
		config_obj: config_obj,
	}, nil
}

func init() {
	result_sets.RegisterTimedResultSetFactory(TimedFactory{})
}
