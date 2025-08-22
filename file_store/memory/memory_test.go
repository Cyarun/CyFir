package memory_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/Cyarun/CyFir/config"
	"github.com/Cyarun/CyFir/file_store"
	"github.com/Cyarun/CyFir/file_store/memory"
	"github.com/Cyarun/CyFir/file_store/tests"
)

type MemoryTestSuite struct {
	*tests.FileStoreTestSuite

	file_store *memory.MemoryFileStore
}

func (self *MemoryTestSuite) SetupTest() {
	self.file_store.Clear()
}

func TestMemoeyFileStore(t *testing.T) {
	config_obj := config.GetDefaultConfig()
	file_store_factory := memory.NewMemoryFileStore(config_obj)

	file_store.OverrideFilestoreImplementation(config_obj, file_store_factory)

	suite.Run(t, &MemoryTestSuite{
		FileStoreTestSuite: tests.NewFileStoreTestSuite(config_obj, file_store_factory),
		file_store:         file_store_factory,
	})
}
