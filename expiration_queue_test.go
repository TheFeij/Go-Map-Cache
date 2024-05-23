package mapcache

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestExpirationQueue_pushEntry(t *testing.T) {
	eq := newExpirationQueue()
	require.NotEmpty(t, eq)

	t.Run("OK", func(t *testing.T) {
		entry := newCacheEntry("key1", "value1", time.Second*10)
		eq.pushEntry(entry)

		poppedEntry, err := eq.popEntry()
		require.NoError(t, err)
		require.Equal(t, entry, poppedEntry)
	})
	t.Run("NilEntry", func(t *testing.T) {
		err := eq.pushEntry(nil)
		require.Error(t, err)
	})
}

func TestExpirationQueue_popEntry(t *testing.T) {
	eq := newExpirationQueue()
	require.NotEmpty(t, eq)

	t.Run("OK", func(t *testing.T) {
		entry1 := newCacheEntry("key1", "value1", time.Hour*10)
		entry2 := newCacheEntry("key2", 2, time.Minute*10)
		eq.pushEntry(entry1)
		eq.pushEntry(entry2)

		poppedEntry, err := eq.popEntry()
		require.NoError(t, err)
		require.Equal(t, entry2, poppedEntry)

		poppedEntry, err = eq.popEntry()
		require.NoError(t, err)
		require.Equal(t, entry1, poppedEntry)

		require.True(t, eq.isEmpty())
	})
	t.Run("EmptyQueue", func(t *testing.T) {
		poppedEntry, err := eq.popEntry()
		require.Error(t, err)
		require.Nil(t, poppedEntry)
	})
}

func TestExpirationQueue_isEarliestExpired(t *testing.T) {
	eq := newExpirationQueue()
	require.NotEmpty(t, eq)

	entry1 := newCacheEntry("key1", nil, -time.Minute)
	entry2 := newCacheEntry("key2", nil, time.Minute)

	err := eq.pushEntry(entry1)
	require.NoError(t, err)
	err = eq.pushEntry(entry2)
	require.NoError(t, err)

	require.True(t, eq.isEarliestExpired())

	_, err = eq.popEntry()
	require.NoError(t, err)

	require.False(t, eq.isEarliestExpired())

	_, err = eq.popEntry()
	require.NoError(t, err)

	require.False(t, eq.isEarliestExpired())
}
