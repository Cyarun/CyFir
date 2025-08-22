package utils

import (
	"github.com/Velocidex/yaml/v2"
	api_proto "github.com/Cyarun/CyFir/api/proto"
	"github.com/Cyarun/CyFir/artifacts/assets"
)

// Loads the api description from the embedded asset
func LoadApiDescription() ([]*api_proto.Completion, error) {
	data, err := assets.ReadFile("/docs/references/vql.yaml")
	if err != nil {
		return nil, err
	}

	result := []*api_proto.Completion{}
	err = yaml.Unmarshal(data, &result)
	return result, err
}
