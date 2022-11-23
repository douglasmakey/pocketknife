package cmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	t.Run("Max int", func(t *testing.T) {
		assert.Equal(t, 15, Max(10, 15))
	})

	t.Run("Max Float", func(t *testing.T) {
		assert.Equal(t, 22.4, Max(22.3, 22.4))
	})

	t.Run("Max Chart", func(t *testing.T) {
		assert.Equal(t, 'a', Max('a', 'A'))
	})
}

func TestMin(t *testing.T) {
	t.Run("Min int", func(t *testing.T) {
		assert.Equal(t, 10, Min(10, 15))
	})

	t.Run("Min Float", func(t *testing.T) {
		assert.Equal(t, 22.3, Min(22.3, 22.4))
	})

	t.Run("Min Chart", func(t *testing.T) {
		assert.Equal(t, 'A', Min('a', 'A'))
	})
}
