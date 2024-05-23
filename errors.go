package mapcache

import (
	"fmt"
)

var (
	ErrKeyNotFound       = fmt.Errorf("key does not exist in the cache")
	ErrCacheUnsuccessful = fmt.Errorf("unable to cache data")
	ErrMemoryLimit       = fmt.Errorf("memory limit exceeded")
	ErrEntrySizeLimit    = fmt.Errorf("entry size limit exceeded")
)
