package memory_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/Cyarun/CyFir/config"
	"github.com/Cyarun/CyFir/file_store"
	"github.com/Cyarun/CyFir/file_store/memory"
	"github.com/Cyarun/CyFir/file_store/tests"

	_ "github.com/Cyarun/CyFir/result_sets/simple"
	_ "github.com/Cyarun/CyFir/result_sets/timed"
)

func TestMemoryQueueManager(t *testing.T) {
	config_obj := config.GetDefaultConfig()
	file_store_factory := memory.NewMemoryFileStore(config_obj)
	manager := memory.NewMemoryQueueManager(config_obj, file_store_factory)

	file_store.OverrideFilestoreImplementation(config_obj, file_store_factory)

	suite.Run(t, tests.NewQueueManagerTestSuite(config_obj, manager, file_store_factory))
}
