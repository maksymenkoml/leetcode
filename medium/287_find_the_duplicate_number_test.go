package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test287(t *testing.T) {
	type examples struct {
		nums []int
		want int
	}
	s := []examples{
		{nums: []int{1, 3, 4, 2, 2}, want: 2},
		{nums: []int{3, 1, 3, 4, 2}, want: 3},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, findDuplicate(tst.nums))
	}
}

func findDuplicate(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]
		if s == f {
			break
		}
	}

	s2 := 0
	for {
		s = nums[s]
		s2 = nums[s2]
		if s == s2 {
			return s
		}
	}
}
