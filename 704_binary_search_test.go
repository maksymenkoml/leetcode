package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test704(t *testing.T) {
	type examples struct {
		s    []int
		t    int
		want int
	}
	s := []examples{
		{s: []int{-1, 0, 3, 5, 9, 12}, t: 9, want: 4},
		{s: []int{-1, 0, 3, 5, 9, 12}, t: 2, want: -1},
	}
	for i := range s {
		require.Equal(t, s[i].want, search(s[i].s, s[i].t), s[i].s)
	}
}

func search(nums []int, target int) int {
	total := len(nums)
	ppos := 0
	for {
		if len(nums) == 0 || (ppos > 0 && total <= ppos+1) {
			return -1
		}
		pos := len(nums) / 2
		n := nums[pos]
		if n == target {
			if ppos > 0 {
				return pos + ppos
			}
			return pos
		}
		if pos == 0 {
			return -1
		}
		if n < target {
			nums = nums[pos:]
			ppos += pos
			continue
		}
		if n > target {
			nums = nums[:pos]
			continue
		}
	}
}
