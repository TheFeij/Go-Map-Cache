package mapcache

import "time"

// CacheService defines an interface for caching functionalities
type CacheService interface {
	// CacheData saves the input key and value to cache
	CacheData(key string, value interface{}, duration time.Duration) error
	// LoadData loads the data relative to the input key
	// returns an error object if key does not exist in the cache
	LoadData(key string) (interface{}, error)
}
