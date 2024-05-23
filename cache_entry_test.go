package mapcache

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCacheEntry(t *testing.T) {
	key := "key1"
	value := "value1"
	duration := time.Second * 10

	entry := newCacheEntry(key, value, duration)
	require.NotEmpty(t, entry)
	require.Equal(t, key, entry.key)
	require.Equal(t, value, entry.value)

	seconds := entry.expiration / 1e9
	nanoseconds := entry.expiration % 1e9
	require.WithinDuration(t, time.Now().Add(duration).UTC(), time.Unix(seconds, nanoseconds).UTC(), time.Millisecond)

	require.False(t, entry.isExpired())
}
