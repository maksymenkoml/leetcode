package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test128(t *testing.T) {
	type examples struct {
		nums []int
		want int
	}
	s := []examples{
		{nums: []int{100, 4, 200, 1, 3, 2}, want: 4},
		{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, want: 9},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, longestConsecutive(tst.nums), tst.nums)
	}
}

func longestConsecutive(nums []int) int {
	store := make(map[int]bool)

	for _, num := range nums {
		store[num] = true
	}

	var maxSeq int
	for num := range store {
		seq := 1
		curSeq := findNext(store, num, seq)
		if curSeq > maxSeq {
			maxSeq = curSeq
		}
	}

	return maxSeq
}

func findNext(nums map[int]bool, find int, seq int) int {
	if _, ok := nums[find+1]; ok {
		seq++
		seq = findNext(nums, find+1, seq)
	}
	return seq
}
