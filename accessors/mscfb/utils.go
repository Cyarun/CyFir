package mscfb

import (
	"github.com/Velocidex/go-mscfb/parser"
	"github.com/Cyarun/CyFir/accessors"
	"github.com/Cyarun/CyFir/constants"
	vql_subsystem "github.com/Cyarun/CyFir/vql"
	"github.com/Cyarun/CyFir/vql/readers"
	"www.velocidex.com/golang/vfilter"
)

func GetMscfbContext(scope vfilter.Scope,
	device, fullpath *accessors.OSPath, accessor string) (
	result *parser.OLEContext, err error) {

	if device == nil {
		device, err = fullpath.Delegate(scope)
		if err != nil {
			return nil, err
		}
		accessor = fullpath.DelegateAccessor()
	}

	return GetMscfbCache(scope, device, accessor)
}

func GetMscfbCache(scope vfilter.Scope,
	device *accessors.OSPath, accessor string) (*parser.OLEContext, error) {
	key := "mscfb_cache" + device.String() + accessor

	// Get the cache context from the root scope's cache
	cache_ctx, ok := vql_subsystem.CacheGet(scope, key).(*parser.OLEContext)
	if !ok {
		lru_size := vql_subsystem.GetIntFromRow(
			scope, scope, constants.NTFS_CACHE_SIZE)

		paged_reader, err := readers.NewAccessorReader(
			scope, accessor, device, int(lru_size))
		if err != nil {
			return nil, err
		}

		cache_ctx, err = parser.GetOLEContext(paged_reader)
		if err != nil {
			return nil, err
		}
		vql_subsystem.CacheSet(scope, key, cache_ctx)

		// Close the device when we are done with this query.
		err = vql_subsystem.GetRootScope(scope).AddDestructor(func() {
			paged_reader.Close()
		})
		if err != nil {
			return nil, err
		}
	}

	return cache_ctx, nil
}
