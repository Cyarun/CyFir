package acls

import (
	"sync"

	acl_proto "github.com/Cyarun/CyFir/acls/proto"
)

var (
	mu             sync.Mutex
	lockdown_token *acl_proto.ApiClientACL
)

func LockdownToken() *acl_proto.ApiClientACL {
	mu.Lock()
	defer mu.Unlock()
	return lockdown_token
}

func SetLockdownToken(token *acl_proto.ApiClientACL) {
	mu.Lock()
	defer mu.Unlock()
	lockdown_token = token
}
