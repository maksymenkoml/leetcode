package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test739(t *testing.T) {
	type examples struct {
		temperatures []int
		want         []int
	}
	s := []examples{
		{temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73}, want: []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{temperatures: []int{30, 40, 50, 60}, want: []int{1, 1, 1, 0}},
		{temperatures: []int{30, 60, 90}, want: []int{1, 1, 0}},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, dailyTemperatures(tst.temperatures), tst.temperatures)
	}
}

func dailyTemperatures(temperatures []int) []int {
	store := make([]int, len(temperatures))

	for i, currTemp := range temperatures {
		found := false
		n := 0
		for k := i + 1; k < len(temperatures); k++ {
			n++
			if temperatures[k] > currTemp {
				store[i] = n
				found = true
				break
			}
		}
		if !found {
			store[i] = 0
		}
	}

	return store
}
