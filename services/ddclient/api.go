package ddclient

import (
	"context"

	config_proto "github.com/Cyarun/CyFir/config/proto"
)

type Updater interface {
	UpdateDDNSRecord(
		ctx context.Context, config_obj *config_proto.Config,
		external_ip string) error
}
