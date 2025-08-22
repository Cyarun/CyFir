package api

import (
	"context"

	"github.com/Cyarun/CyFir/api/proto"
	api_proto "github.com/Cyarun/CyFir/api/proto"
)

func (self *ApiServer) Check(
	ctx context.Context,
	in *api_proto.HealthCheckRequest) (*api_proto.HealthCheckResponse, error) {

	return &proto.HealthCheckResponse{
		Status: api_proto.HealthCheckResponse_SERVING,
	}, nil
}
