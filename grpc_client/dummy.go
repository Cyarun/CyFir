package grpc_client

import (
	"context"
	"sync"

	api_proto "github.com/Cyarun/CyFir/api/proto"
	config_proto "github.com/Cyarun/CyFir/config/proto"
)

type DummyGRPCAPIClient struct {
	mu sync.Mutex
	*GRPCAPIClient
}

func (self *DummyGRPCAPIClient) GetAPIClient(
	ctx context.Context,
	identity CallerIdentity,
	config_obj *config_proto.Config) (
	api_proto.APIClient, func() error, error) {
	self.mu.Lock()
	client := self.GRPCAPIClient
	if client == nil {
		new_client, err := NewGRPCAPIClient(config_obj)
		if err != nil {
			return nil, nil, err
		}
		self.GRPCAPIClient = new_client
		client = new_client
	}
	self.mu.Unlock()

	return client.GetAPIClient(ctx, identity, config_obj)
}
