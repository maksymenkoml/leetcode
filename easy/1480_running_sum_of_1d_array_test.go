package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1480(t *testing.T) {
	type examples struct {
		nums []int
		want []int
	}
	s := []examples{
		{nums: []int{1, 2, 3, 4}, want: []int{1, 3, 6, 10}},
		{nums: []int{1, 1, 1, 1, 1}, want: []int{1, 2, 3, 4, 5}},
		{nums: []int{3, 1, 2, 10, 1}, want: []int{3, 4, 6, 16, 17}},
	}
	for _, tst := range s {
		got := runningSum(tst.nums)

		require.Equal(t, tst.want, got)
	}
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
