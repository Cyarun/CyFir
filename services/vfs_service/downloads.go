package vfs_service

import (
	"context"

	config_proto "github.com/Cyarun/CyFir/config/proto"
	"github.com/Cyarun/CyFir/file_store"
	flows_proto "github.com/Cyarun/CyFir/flows/proto"
	"github.com/Cyarun/CyFir/json"
	"github.com/Cyarun/CyFir/paths"
	"github.com/Cyarun/CyFir/result_sets"
	"github.com/Cyarun/CyFir/utils"
)

func (self *VFSService) WriteDownloadInfo(
	ctx context.Context,
	config_obj *config_proto.Config,
	client_id string,
	accessor string,
	client_components []string,
	record *flows_proto.VFSDownloadInfo) error {

	// We are only interested in the directory that the file in
	// contained in.
	components := append([]string{accessor}, client_components...)
	dir_components := components[:len(components)-1]
	record.Name = components[len(components)-1]

	client_path_manager := paths.NewClientPathManager(client_id)
	file_store_factory := file_store.GetFileStore(config_obj)
	writer, err := result_sets.NewResultSetWriter(
		file_store_factory,
		client_path_manager.VFSDownloadInfoResultSet(dir_components),
		json.DefaultEncOpts(),
		utils.BackgroundWriter,
		result_sets.AppendMode)
	if err != nil {
		return err
	}
	defer writer.Close()

	serialized, err := json.Marshal(record)
	if err != nil {
		return err
	}
	serialized = append(serialized, '\n')
	writer.WriteJSONL(serialized, 1)

	return nil
}
