package main

import "sync"

// TypedBinding is a generic binding which introduces generics to
// Fyne's builtin bindings.
//
// It shares the listener functionality, but builds on top of it by passing
// the changed value to the listener.
type TypedBinding[T any] struct {
	data      T
	listeners []func(T)
	m         sync.RWMutex
}

// NewTypedBinding creates a new binding for the type T.
func NewTypedBinding[T any]() *TypedBinding[T] {
	return &TypedBinding[T]{}
}

// Set updates the internal representation of the data.
func (b *TypedBinding[T]) Set(data T) {
	b.m.Lock()
	b.data = data
	b.m.Unlock()
	for _, listener := range b.listeners {
		listener(b.data)
	}
}

// Get returns the internal representation of the data.
func (b *TypedBinding[T]) Get() T {
	b.m.RLock()
	defer b.m.RUnlock()
	return b.data
}

// AddListener adds a listener which is called whenever the internal representation is changed
// via Set.
func (b *TypedBinding[T]) AddListener(listener func(data T)) {
	b.listeners = append(b.listeners, listener)
}
