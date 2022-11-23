package datatype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		var stack Stack[int]

		assert.True(t, stack.IsEmpty())
		stack.Push(1)
		assert.False(t, stack.IsEmpty())
		stack.Push(2)
		value, _ := stack.Pop()
		assert.Equal(t, value, 2)
		value, _ = stack.Pop()
		assert.Equal(t, value, 1)
		assert.True(t, stack.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {
		var stack Stack[string]
		assert.True(t, stack.IsEmpty())
		stack.Push("A")
		assert.False(t, stack.IsEmpty())
		stack.Push("B")
		value, _ := stack.Pop()
		assert.Equal(t, value, "B")
		value, _ = stack.Pop()
		assert.Equal(t, value, "A")
		assert.True(t, stack.IsEmpty())
	})

	t.Run("custom stack", func(t *testing.T) {
		type custom string
		var stack Stack[custom]
		assert.True(t, stack.IsEmpty())
		stack.Push("A")
		assert.False(t, stack.IsEmpty())
		stack.Push("B")
		value, _ := stack.Pop()
		assert.Equal(t, value, custom("B"))
		value, _ = stack.Pop()
		assert.Equal(t, value, custom("A"))
		assert.True(t, stack.IsEmpty())
	})
}
