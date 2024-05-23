// Package mapcache provides a simple in-memory cache implementation using a map-based storage.
//
// # Overview
//
// This package offers a basic caching mechanism that allows storing key-value pairs with expiration times.
// It includes functionality to set a maximum memory limit and maximum entry size for the cache.
// The cache is implemented using a map to store key-value pairs, and a background goroutine is used
// to periodically delete expired entries from the cache.
//
// # Usage
//
// To use the map_cache package, you typically create a singleton instance of MapCache using the GetMapCache function,
// passing in parameters such as the maximum memory limit, maximum entry size, and cleanup duration.
// You can then use the CacheData method to store data in the cache and the LoadData method to retrieve data from the cache.
//
// Example:
//
//	cache := GetMapCache(1000, 100, time.Minute)
//	err := cache.CacheData("key", "value", time.Second * 30)
//	if err != nil {
//	    fmt.Println("Error caching data:", err)
//	}
//
//	data, err := cache.LoadData("key")
//	if err != nil {
//	    fmt.Println("Error loading data:", err)
//	} else {
//	    fmt.Println("Data from cache:", data)
//	}
//
// # Errors
//
// The package defines several common errors that can occur during cache operations,
// such as key not found, cache operation failure, memory limit exceeded, and entry size limit exceeded.
// These errors are exposed as package-level variables and can be used for error handling in client code.
//
//	ErrKeyNotFound       = fmt.Errorf("key does not exist in the cache")
//	ErrCacheUnsuccessful = fmt.Errorf("unable to cache data")
//	ErrMemoryLimit       = fmt.Errorf("memory limit exceeded")
//	ErrEntrySizeLimit    = fmt.Errorf("entry size limit exceeded")
//
// # Thread Safety
//
// The MapCache implementation in this package is designed to be goroutine-safe.
// It uses sync.RWMutex to handle concurrent access safely during cache operations.
// However, it's essential to note that concurrent access can impact performance,
// especially under heavy load or contention.
//
// # Notes
//
//   - This package is primarily intended for use cases where a simple, in-memory cache is sufficient.
//     For more advanced caching needs, consider using dedicated caching solutions such as Redis or Memcached.
//   - Be mindful of the memory usage when setting the maximum memory limit and entry size limit,
//     as exceeding these limits can lead to performance issues or out-of-memory errors.
package mapcache
