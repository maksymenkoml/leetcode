package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test977(t *testing.T) {
	type examples struct {
		nums []int
		want []int
	}
	s := []examples{
		{nums: []int{-2, 3}, want: []int{4, 9}},
		{nums: []int{-7, -3, 2, 3, 11}, want: []int{4, 9, 9, 49, 121}},
		{nums: []int{-4, -1, 0, 3, 10}, want: []int{0, 1, 9, 16, 100}},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, sortedSquares(tst.nums), tst.nums)
	}
}

func sortedSquares(nums []int) []int {
	var res []int
	L, R := 0, len(nums)-1
	for {
		if abs(nums[L]) > abs(nums[R]) {
			res = append(res, []int{nums[L] * nums[L]}...)
			L++
		} else {
			res = append(res, []int{nums[R] * nums[R]}...)
			R--
		}
		if L > R {
			arrRev(res)
			return res
		}
	}
}

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

func arrRev(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
