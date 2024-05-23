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

// GetMapCache initializes and returns an instance of MapCache
func GetMapCache(maxMemory int, maxEntrySize int, cleanupDuration time.Duration) *MapCache {
	mapCacheInstance := &MapCache{
		storage:      make(map[string]*mapCacheEntry),
		expirationQ:  newExpirationQueue(),
		maxMemory:    int64(maxMemory),
		maxEntrySize: int64(maxEntrySize),
	}

	// Start a background goroutine to delete expired entries periodically
	go mapCacheInstance.cleanupExpiredEntries(cleanupDuration)

	return mapCacheInstance
}

// cleanupExpiredEntries periodically checks for and deletes expired entries
func (c *MapCache) cleanupExpiredEntries(cleanupDuration time.Duration) {
	for {
		time.Sleep(cleanupDuration)

		c.rwMutex.Lock()
		for {
			// if expiration
			if c.expirationQ.isEmpty() || !c.expirationQ.isEarliestExpired() {
				break
			}

			// Pop the expired entry from the expiration queue
			// error is discarded because it is checked earlier that queue is not empty
			entry, _ := c.expirationQ.popEntry()

			// Remove the entry from the storage map
			c.deleteEntry(entry)
		}
		c.rwMutex.Unlock()
	}
}
