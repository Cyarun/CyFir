package utils

import actions_proto "github.com/Cyarun/CyFir/actions/proto"

func GetQueryName(args []*actions_proto.VQLRequest) string {
	for _, query := range args {
		if query.Name != "" {
			return query.Name
		}
	}
	return ""
}
