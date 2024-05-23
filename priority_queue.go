package mapcache

// priorityQueue implements a priority queue for expired entries.
// implements container/heap interface
type priorityQueue []*mapCacheEntry

// Len returns the number of elements in the queue.
func (pq priorityQueue) Len() int { return len(pq) }

// Less reports whether the element with index i should sort before the element with index j.
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].expiration < pq[j].expiration
}

// Swap swaps the elements with indexes i and j.
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push pushes the element x onto the queue.
func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*mapCacheEntry)
	*pq = append(*pq, item)
}

// Pop removes and returns the minimum element (according to Less) from the queue.
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // Avoid memory leak
	*pq = old[0 : n-1]
	return item
}
