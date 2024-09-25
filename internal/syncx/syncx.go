// Package syncx holds custom structus for parallelization.
package syncx

import (
	"sync"
)

// Map wraps Go's builtin map with a mutex.
//
// Unlike mutexes, but like maps, Map can be copied after initialization.
type Map[K comparable, V any] struct {
	mp    map[K]V
	mutex *sync.RWMutex
}

// NewMap creates and returns a new map instance.
func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{
		mp:    make(map[K]V),
		mutex: new(sync.RWMutex),
	}
}

// Len is like len(map).
func (m Map[K, V]) Len() int {
	m.mutex.RLock()
	n := len(m.mp)
	m.mutex.RUnlock()
	return n
}

// Get is like map[k].
func (m Map[K, V]) Get(k K) (V, bool) {
	m.mutex.RLock()
	v, ok := m.mp[k]
	m.mutex.RUnlock()
	return v, ok
}

// Set is like map[k] = v.
func (m Map[K, V]) Set(k K, v V) {
	m.mutex.Lock()
	m.mp[k] = v
	m.mutex.Unlock()
}

// Delete is like delete(map, k).
func (m Map[K, V]) Delete(k K) {
	m.mutex.Lock()
	delete(m.mp, k)
	m.mutex.Unlock()
}

// Pop returns the value and deletes it afterwards.
// This should be preferred over calling Get() and Delete() sequentially to avoid additional mutex locks.
func (m Map[K, V]) Pop(k K) (V, bool) {
	m.mutex.Lock()
	v, ok := m.mp[k]
	delete(m.mp, k)
	m.mutex.Unlock()
	return v, ok
}
