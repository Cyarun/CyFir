package sanity

import (
	"context"

	config_proto "github.com/Cyarun/CyFir/config/proto"
	"github.com/Cyarun/CyFir/datastore"
	"github.com/Cyarun/CyFir/file_store/api"
	"github.com/Cyarun/CyFir/logging"
	"github.com/Cyarun/CyFir/paths"
	"github.com/Cyarun/CyFir/services"
)

func maybeMigrateClientIndex(
	ctx context.Context, config_obj *config_proto.Config) error {

	db, err := datastore.GetDB(config_obj)
	if err != nil {
		return err
	}

	items, err := db.ListChildren(config_obj, paths.CLIENT_INDEX_URN)
	if err != nil {
		return err
	}

	if len(items) > 0 {
		return nil
	}

	logger := logging.GetLogger(config_obj, &logging.FrontendComponent)
	logger.Info("Converting legacy client index to new format")

	count := 0

	indexer, err := services.GetIndexer(config_obj)
	if err != nil {
		return err
	}

	// Migrate the old index to the new index.
	err = datastore.Walk(config_obj, db, paths.CLIENT_INDEX_URN_DEPRECATED,
		datastore.WalkWithoutDirectories,
		func(path api.DSPathSpec) error {
			client_id := path.Base()
			term := path.Dir().Base()
			count++
			if count%500 == 0 {
				logger.Info("Converted %v index items to the new format", count)
			}
			return indexer.SetIndex(client_id, term)
		})

	return err
}
