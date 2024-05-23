package mapcache

import (
	"fmt"
	"time"
)

// CacheData saves the input key and value to cache
func (c *MapCache) CacheData(key string, value interface{}, duration time.Duration) (err error) {
	// capture panics
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			err = fmt.Errorf("%w: %v", ErrCacheUnsuccessful, r)
		}
	}()

	// create the new entry
	newEntry := newCacheEntry(key, value, duration)

	// acquire read lock and then release it at the end of the function
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	// Check if the entry size exceeds the max entry size
	if newEntry.size > c.maxEntrySize {
		err = ErrEntrySizeLimit
		return
	}

	// Check if the new entry exceeds the memory limit
	if c.curMemory+newEntry.size > c.maxMemory {
		err = ErrMemoryLimit
		return
	}

	// store key and value in the cache
	c.storage[key] = newEntry

	return
}

// LoadData loads the data relative to the input key
// returns an error object if key does not exist in the cache
func (c *MapCache) LoadData(key string) (value interface{}, err error) {
	// acquire read lock and then release it at the end of the function
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	cacheEntry, found := c.storage[key]
	if !found || cacheEntry.isExpired() {
		err = ErrKeyNotFound
		return
	}

	return cacheEntry.value, nil
}

// deleteEntry Removes the entry from cache
func (c *MapCache) deleteEntry(entry *mapCacheEntry) {
	delete(c.storage, entry.key)
	c.curMemory -= entry.size
}
