package test_utils

import (
	"archive/zip"
	"io"

	config_proto "github.com/Cyarun/CyFir/config/proto"
	"github.com/Cyarun/CyFir/file_store"
	"github.com/Cyarun/CyFir/file_store/api"
	"github.com/Cyarun/CyFir/utils"
)

func UnzipToFilestore(
	config_obj *config_proto.Config,
	base api.FSPathSpec,
	zip_path string) error {

	reader, err := zip.OpenReader(zip_path)
	if err != nil {
		return err
	}
	defer reader.Close()

	file_store_factory := file_store.GetFileStore(config_obj)

	for _, f := range reader.File {
		components := utils.SplitComponents(f.Name)
		output_path := base.AddChild(components...).
			SetType(api.PATH_TYPE_FILESTORE_ANY)
		fd, err := file_store_factory.WriteFile(output_path)
		if err != nil {
			return err
		}
		infd, err := reader.Open(f.Name)
		if err != nil {
			continue
		}

		_, err = io.Copy(fd, infd)
		fd.Close()

		if err != nil {
			return err
		}
	}

	return nil
}
