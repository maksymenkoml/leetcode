package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test283(t *testing.T) {
	type examples struct {
		nums []int
		want []int
	}
	s := []examples{
		{nums: []int{1, 1, 0, 1}, want: []int{1, 1, 1, 0}},
		{nums: []int{0, 0, 1, 0, 1, 2}, want: []int{1, 1, 2, 0, 0, 0}},
		{nums: []int{2, 1}, want: []int{2, 1}},
		{nums: []int{1, 0, 1}, want: []int{1, 1, 0}},
		{nums: []int{1}, want: []int{1}},
		{nums: []int{0, 0, 1}, want: []int{1, 0, 0}},
		{nums: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		{nums: []int{0}, want: []int{0}},
	}
	for _, tst := range s {
		moveZeroes(tst.nums)
		require.Equal(t, tst.want, tst.nums, tst.nums)
	}
}

func moveZeroes(nums []int) {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		}
	}
}

func moveZeroes1(nums []int) {
	if len(nums) == 1 {
		return
	}

	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[p] == 0 {
			nums = append(nums[:p], append(nums[p+1:], 0)...)
			continue
		}
		p++
	}

	return
}
