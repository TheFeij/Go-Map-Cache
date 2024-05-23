package mapcache

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMapCache_CacheData(t *testing.T) {
	mapCache := NewMapCache(4096, 128, time.Second)

	t.Run("OK", func(t *testing.T) {
		mapCache.CacheData("key1", "key1", 10*time.Second)

		value1, err := mapCache.LoadData("key1")
		require.Equal(t, "key1", value1)
		require.NoError(t, err)
	})
	t.Run("InvalidCache", func(t *testing.T) {
		// using an invalid cache, cache should be gotten from GetMapCache
		invalidCache := MapCache{}
		err := invalidCache.CacheData("key1", "value1", time.Hour)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrEntrySizeLimit)
	})
	t.Run("MemoryLimitExceeded", func(t *testing.T) {
		smallCache := NewMapCache(4, 8, time.Hour)
		err := smallCache.CacheData("key1", "value1", time.Minute)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrMemoryLimit)
	})
}

func TestMapCache_LoadData(t *testing.T) {
	mapCache := NewMapCache(4096, 128, time.Second)

	t.Run("OK", func(t *testing.T) {
		mapCache.CacheData("key1", "key1", 10*time.Second)

		value1, err := mapCache.LoadData("key1")
		require.Equal(t, "key1", value1)
		require.NoError(t, err)
	})
	t.Run("NonExistingKey", func(t *testing.T) {
		// trying to get the value of a non-existent key
		_, err := mapCache.LoadData("Non Existence key")
		require.Error(t, err)
		require.Equal(t, err, ErrKeyNotFound)
	})
}
