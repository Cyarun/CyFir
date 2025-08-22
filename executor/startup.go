package executor

import (
	"context"
	"sync"

	config_proto "github.com/Cyarun/CyFir/config/proto"
	"github.com/Cyarun/CyFir/logging"
)

func RunStartupTasks(
	ctx context.Context,
	config_obj *config_proto.Config,
	wg *sync.WaitGroup,
	exe Executor) error {

	err := CheckForCrashes(ctx, config_obj, wg, exe)
	if err != nil {
		// Not a fatal error, just move on
		logger := logging.GetLogger(config_obj, &logging.ClientComponent)
		logger.Error("<red>CheckForCrashes Error:</> %v", err)
	}

	return nil
}
