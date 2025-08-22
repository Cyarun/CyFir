package fat

import (
	fat "github.com/Velocidex/go-fat/parser"
	"github.com/Cyarun/CyFir/accessors"
	"github.com/Cyarun/CyFir/constants"
	vql_subsystem "github.com/Cyarun/CyFir/vql"
	"github.com/Cyarun/CyFir/vql/readers"
	"www.velocidex.com/golang/vfilter"
)

func GetFatContext(scope vfilter.Scope,
	device, fullpath *accessors.OSPath, accessor string) (
	result *fat.FATContext, err error) {

	if device == nil {
		device, err = fullpath.Delegate(scope)
		if err != nil {
			return nil, err
		}
		accessor = fullpath.DelegateAccessor()
	}

	return GetFATCache(scope, device, accessor)
}

func GetFATCache(scope vfilter.Scope,
	device *accessors.OSPath, accessor string) (*fat.FATContext, error) {
	key := "fat_cache" + device.String() + accessor

	// Get the cache context from the root scope's cache
	cache_ctx, ok := vql_subsystem.CacheGet(scope, key).(*fat.FATContext)
	if !ok {
		lru_size := vql_subsystem.GetIntFromRow(
			scope, scope, constants.NTFS_CACHE_SIZE)

		paged_reader, err := readers.NewAccessorReader(
			scope, accessor, device, int(lru_size))
		if err != nil {
			return nil, err
		}

		cache_ctx, err = fat.GetFATContext(paged_reader)
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
