package syncx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert := assert.New(t)

	m := NewMap[string, int]()

	_, exists := m.Get("Foo")
	assert.False(exists)

	m.Set("Foo", 999)
	out, exists := m.Get("Foo")
	assert.True(exists)
	assert.Equal(999, out)

	m.Delete("Foo")

	_, exists = m.Get("Foo")
	assert.False(exists)
}
