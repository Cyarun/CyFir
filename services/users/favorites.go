package users

import (
	"context"

	api_proto "github.com/Cyarun/CyFir/api/proto"
	config_proto "github.com/Cyarun/CyFir/config/proto"
)

func (self UserManager) GetFavorites(
	ctx context.Context,
	config_obj *config_proto.Config,
	principal, fav_type string) (*api_proto.Favorites, error) {
	return self.storage.GetFavorites(ctx, config_obj, principal, fav_type)
}
