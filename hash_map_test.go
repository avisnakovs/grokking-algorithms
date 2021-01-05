package grokking_algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHashMap(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		m := NewHashMap()
		value, ok := m.Get("")
		assert.Empty(t, value)
		assert.False(t, ok)
	})

	t.Run("set get map", func(t *testing.T) {
		m := NewHashMap()
		m.Set("key", "value")
		value, ok := m.Get("key")
		assert.Equal(t, "value", value)
		assert.True(t, ok)
	})

	t.Run("increase map size", func(t *testing.T) {
		m := NewHashMap()
		m.Set("1", "1")
		m.Set("2", "2")
		m.Set("3", "3")
		m.Set("4", "4")
		m.Set("5", "5")
		m.Set("6", "6")
		m.Set("7", "7")
		m.Set("8", "8")
		m.Set("9", "9")
		m.Set("10", "10")
		m.Set("11", "11")

		value, ok := m.Get("1")
		assert.Equal(t, "1", value)
		assert.True(t, ok)

		value, ok = m.Get("2")
		assert.Equal(t, "2", value)
		assert.True(t, ok)

		value, ok = m.Get("3")
		assert.Equal(t, "3", value)
		assert.True(t, ok)

		value, ok = m.Get("4")
		assert.Equal(t, "4", value)
		assert.True(t, ok)

		value, ok = m.Get("5")
		assert.Equal(t, "5", value)
		assert.True(t, ok)

		value, ok = m.Get("6")
		assert.Equal(t, "6", value)
		assert.True(t, ok)

		value, ok = m.Get("7")
		assert.Equal(t, "7", value)
		assert.True(t, ok)

		value, ok = m.Get("8")
		assert.Equal(t, "8", value)
		assert.True(t, ok)

		value, ok = m.Get("9")
		assert.Equal(t, "9", value)
		assert.True(t, ok)

		value, ok = m.Get("10")
		assert.Equal(t, "10", value)
		assert.True(t, ok)

		value, ok = m.Get("11")

		assert.Equal(t, "11", value)
		assert.True(t, ok)

		value, ok = m.Get("12")
		assert.Empty(t, value)
		assert.False(t, ok)

		hashMap, ok := m.(*hashMap)
		assert.True(t, ok)
		assert.Equal(t, 11, hashMap.size)
		assert.Equal(t, uint64(20), hashMap.len)
	})
}
