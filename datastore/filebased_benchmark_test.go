package datastore_test

import (
	"os"
	"testing"

	actions_proto "github.com/Cyarun/CyFir/actions/proto"
	"github.com/Cyarun/CyFir/config"
	"github.com/Cyarun/CyFir/datastore"
	"github.com/Cyarun/CyFir/paths"
	"github.com/Cyarun/CyFir/utils/tempfile"
)

func BenchmarkSetSubject(b *testing.B) {
	dir, _ := tempfile.TempDir("datastore_test")
	defer os.RemoveAll(dir) // clean up

	config_obj := config.GetDefaultConfig()
	config_obj.Datastore.FilestoreDirectory = dir
	config_obj.Datastore.Location = dir

	db, _ := datastore.GetDB(config_obj)

	client_id := "C.1234"
	client_path_manager := paths.NewClientPathManager(client_id)

	b.Run("SetSubject", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			client_info := &actions_proto.ClientInfo{
				ClientId: client_id,
			}
			db.SetSubject(config_obj,
				client_path_manager.Ping(), client_info)
		}
	})

}
