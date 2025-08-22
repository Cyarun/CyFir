package paths

import (
	"github.com/Cyarun/CyFir/file_store/api"
	"github.com/Cyarun/CyFir/file_store/path_specs"
)

func DSPathSpecFromClientPath(client_path string) api.DSPathSpec {
	components := ExtractClientPathComponents(client_path)
	result := path_specs.NewUnsafeDatastorePath(components...)
	if len(components) > 0 {
		last := len(components) - 1
		name_type, name := api.GetDataStorePathTypeFromExtension(
			components[last])
		components[last] = name
		return result.SetType(name_type)
	}
	return result
}

func FSPathSpecFromClientPath(client_path string) api.FSPathSpec {
	components := ExtractClientPathComponents(client_path)
	result := path_specs.NewUnsafeFilestorePath(components...)
	if len(components) > 0 {
		last := len(components) - 1
		name_type, name := api.GetFileStorePathTypeFromExtension(
			components[last])
		components[last] = name
		return result.SetType(name_type)
	}
	return result
}
