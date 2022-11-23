package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	type args[T any] struct {
		s []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}

	t.Run("reverse int", func(t *testing.T) {
		tests := []testCase[int]{
			{
				name: "some ints",
				args: args[int]{s: []int{1, 2, 3, 4, 5}},
				want: []int{5, 4, 3, 2, 1},
			},
			{
				name: "some ints",
				args: args[int]{s: []int{10, 15, 20, 33, 44}},
				want: []int{44, 33, 20, 15, 10},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				Reverse(tt.args.s)
			})

			assert.Equal(t, tt.want, tt.args.s)
		}
	})

	t.Run("reverse string", func(t *testing.T) {

		tests := []testCase[string]{
			{
				name: "Case 1",
				args: args[string]{s: []string{"Hello", "world"}},
				want: []string{"world", "Hello"},
			},
			{
				name: "Case 2",
				args: args[string]{s: []string{"a", "b", "c", "d"}},
				want: []string{"d", "c", "b", "a"},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				Reverse(tt.args.s)
			})

			assert.Equal(t, tt.want, tt.args.s)
		}
	})
}
