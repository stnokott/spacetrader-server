// Package mocks provides mock struct instances for testing.
package mocks

import "fmt"

func ptr[T any](x T) *T {
	return &x
}

// Chunk splits src into chunks of size chunkSize and returns the chunk at index chunkIndex.
func Chunk[T any](src []T, chunkIndex int, chunkSize int) []T {
	if chunkIndex*chunkSize >= len(src) {
		panic(fmt.Sprintf("chunk index %d out of bounds for len(src)=%d and chunkSize=%d", chunkIndex, len(src), chunkSize))
	}
	start := chunkIndex * chunkSize
	end := start + chunkSize
	if end > len(src) {
		end = len(src)
	}
	return src[start:end]
}
