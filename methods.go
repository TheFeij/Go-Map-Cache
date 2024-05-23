package mapcache

// deleteEntry Removes the entry from cache
func (c *MapCache) deleteEntry(entry *mapCacheEntry) {
	delete(c.storage, entry.key)
	c.curMemory -= entry.size
}
