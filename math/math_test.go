package math

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestAbs(t *testing.T) {
	type args[T constraints.Integer | constraints.Float] struct {
		a T
	}

	type testCase[T constraints.Integer | constraints.Float] struct {
		name string
		args args[T]
		want T
	}

	t.Run("Abs int", func(t *testing.T) {
		tests := []testCase[int]{
			{
				name: "15",
				args: args[int]{a: 15},
				want: 15,
			},
			{
				name: "-15",
				args: args[int]{a: -15},
				want: 15,
			},
			{
				name: "0",
				args: args[int]{a: 0},
				want: 0,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Abs(tt.args.a); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Abs() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Abs float", func(t *testing.T) {
		tests := []testCase[float32]{
			{
				name: "15.3",
				args: args[float32]{a: 15.3},
				want: 15.3,
			},
			{
				name: "-15.5",
				args: args[float32]{a: -15.5},
				want: 15.5,
			},
			{
				name: "0",
				args: args[float32]{a: 0},
				want: 0,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Abs(tt.args.a); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Abs() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
