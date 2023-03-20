package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test189(t *testing.T) {
	type examples struct {
		nums []int
		k    int
		want []int
	}
	s := []examples{
		{nums: []int{-1}, k: 2, want: []int{-1}},
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, want: []int{5, 6, 7, 1, 2, 3, 4}},
		{nums: []int{-1, -100, 3, 99}, k: 2, want: []int{3, 99, -1, -100}},
	}
	for _, tst := range s {
		rotate(tst.nums, tst.k)

		require.Equal(t, tst.want, tst.nums)
	}
}

func rotate(nums []int, k int) {
	sliceIntReverse(nums)
	kmod := k % len(nums)
	sliceIntReverse(nums[:kmod])
	sliceIntReverse(nums[k%len(nums):])
}

func sliceIntReverse(in []int) {
	l, r := 0, len(in)-1
	for l < r {
		in[l], in[r] = in[r], in[l]
		l++
		r--
	}
}
