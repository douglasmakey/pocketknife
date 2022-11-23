package slice

import (
	"sort"

	"github.com/douglasmakey/pocketknife/math"
	"golang.org/x/exp/constraints"
)

func Reverse[T ~[]K, K any](s T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Sum[T constraints.Integer | constraints.Float](s []T) T {
	var total int
	for n := range s {
		total += n
	}
	return T(total)
}

func SortIntsByAbs(s []int) {
	sort.Slice(s, func(i, j int) bool {
		return math.Abs(i) < math.Abs(j)
	})
}
