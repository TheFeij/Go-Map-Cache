package mapcache

import (
	"container/heap"
	"fmt"
)

// expirationQueue is a priority queue for entries based on their expiration time
// such that the highest priority entry is the entry with the earliest expiration time
//
// it is a wrapper for type priorityQueue to provide easier api
type expirationQueue struct {
	queue priorityQueue
}

// newExpirationQueue returns a new expirationQueue instance
func newExpirationQueue() *expirationQueue {
	return &expirationQueue{queue: priorityQueue{}}
}

// isEarliestExpired checks if the entry with the earliest expiration time is expired or not
// in other words it says if there are any expired entries in the queue
func (eq *expirationQueue) isEarliestExpired() bool {
	if eq.isEmpty() {
		return false
	}
	return eq.queue[0].isExpired()
}

// checks if expirationQueue is empty or not
func (eq *expirationQueue) isEmpty() bool {
	return len(eq.queue) == 0
}

// pop pops the earliest expired entry
func (eq *expirationQueue) popEntry() (*mapCacheEntry, error) {
	if eq.isEmpty() {
		return nil, fmt.Errorf("empty queue")
	}

	return heap.Pop(&eq.queue).(*mapCacheEntry), nil
}

// pushes the entry into the expiration queue
func (eq *expirationQueue) pushEntry(entry *mapCacheEntry) (err error) {
	if entry == nil {
		err = fmt.Errorf("nil entry")
	}

	heap.Push(&eq.queue, entry)
	return
}
