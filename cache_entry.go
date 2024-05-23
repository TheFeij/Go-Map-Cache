package mapcache

import (
	"time"
)

// mapCacheEntry represents a value stored in the cache with an expiration time
type mapCacheEntry struct {
	key        string      // key of the entry
	value      interface{} // value of the entry
	expiration int64       // expiration time of the entry (unix nano)
	size       int64       // size of the entry in bytes
}

// newCacheEntry initializes and returns a mapCacheEntry instance
func newCacheEntry(key string, value interface{}, duration time.Duration) *mapCacheEntry {
	return &mapCacheEntry{
		key:        key,
		value:      value,
		expiration: time.Now().Add(duration).UTC().UnixNano(),
		size:       calculateSize(value),
	}
}

// isExpired checks if entry is expired or not
func (e mapCacheEntry) isExpired() bool {
	return e.expiration < time.Now().UTC().UnixNano()
}
