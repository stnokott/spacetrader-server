// Package syncx holds custom structus for parallelization.
package syncx

import (
	"github.com/ietxaniz/delock"
	log "github.com/sirupsen/logrus"
)

// Map wraps Go's builtin map with a mutex.
type Map[K comparable, V any] struct {
	mp    map[K]V
	mutex delock.RWMutex
}

// NewMap creates and returns a new map instance.
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		mp: make(map[K]V),
	}
}

// Len is like len(map).
func (m *Map[K, V]) Len() int {
	id, err := m.mutex.RLock()
	if err != nil {
		log.Fatalf("deadlock: %v", err)
	}
	n := len(m.mp)
	m.mutex.RUnlock(id)
	return n
}

// Get is like map[k].
func (m *Map[K, V]) Get(k K) (V, bool) {
	id, err := m.mutex.RLock()
	if err != nil {
		log.Fatalf("deadlock: %v", err)
	}
	v, ok := m.mp[k]
	m.mutex.RUnlock(id)
	return v, ok
}

// Set is like map[k] = v.
func (m *Map[K, V]) Set(k K, v V) {
	id, err := m.mutex.Lock()
	if err != nil {
		log.Fatalf("deadlock: %v", err)
	}
	m.mp[k] = v
	m.mutex.Unlock(id)
}

// Delete is like delete(map, k).
func (m *Map[K, V]) Delete(k K) {
	id, err := m.mutex.Lock()
	if err != nil {
		log.Fatalf("deadlock: %v", err)
	}
	delete(m.mp, k)
	m.mutex.Unlock(id)
}

// Pop returns the value and deletes it afterwards.
// This should be preferred over calling Get() and Delete() sequentially to avoid additional mutex locks.
func (m *Map[K, V]) Pop(k K) (V, bool) {
	id, err := m.mutex.Lock()
	if err != nil {
		log.Fatalf("deadlock: %v", err)
	}
	v, ok := m.mp[k]
	delete(m.mp, k)
	m.mutex.Unlock(id)
	return v, ok
}
