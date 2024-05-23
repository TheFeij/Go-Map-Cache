package mapcache

import (
	"sync"
	"time"
)

// MapCache implements the CacheService interface
// stores data in a map with expiration times
type MapCache struct {
	storage      map[string]*mapCacheEntry // map to store key value pairs
	expirationQ  *expirationQueue          // priority queue for expired entries
	rwMutex      sync.RWMutex              // mutex to handle concurrent access safely
	maxMemory    int64                     // maximum memory limit in bytes
	curMemory    int64                     // current memory usage in bytes
	maxEntrySize int64                     // maximum size of each entry
}

// ensures MapCache implements the CacheService interface
var _ CacheService = (*MapCache)(nil)

// mapCacheInstance is the singleton instance of MapCache
var mapCacheInstance *MapCache
var once sync.Once

// GetMapCache returns the singleton instance of MapCache
func GetMapCache(maxMemory int, maxEntrySize int, cleanupDuration time.Duration) *MapCache {
	once.Do(func() {
		mapCacheInstance = &MapCache{
			storage:      make(map[string]*mapCacheEntry),
			expirationQ:  newExpirationQueue(),
			maxMemory:    int64(maxMemory),
			maxEntrySize: int64(maxEntrySize),
		}

		// TODO: Start a background goroutine to delete expired entries periodically
	})
	return mapCacheInstance
}
