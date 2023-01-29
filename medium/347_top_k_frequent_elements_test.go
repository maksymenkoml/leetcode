package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test347(t *testing.T) {
	type examples struct {
		nums []int
		k    int
		want []int
	}
	s := []examples{
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{nums: []int{1}, k: 1, want: []int{1}},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, topKFrequent(tst.nums, tst.k), tst.nums)
	}
}

func topKFrequent(nums []int, k int) []int {
	store := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, num := range nums {
		store[num]++
	}

	for num, count := range store {
		freq[count] = append(freq[count], num)
	}

	var res []int
	for i := len(freq) - 1; i > 0; i-- {
		if freq[i] == nil {
			continue
		}

		for _, num := range freq[i] {
			res = append(res, num)
			k--

			if k == 0 {
				return res
			}
		}
	}
	return nil
}
