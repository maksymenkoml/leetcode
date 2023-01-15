package easy

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test217(t *testing.T) {
	type examples struct {
		s    []int
		want bool
	}
	s := []examples{
		{s: []int{1, 2, 3, 1}, want: true},
		{s: []int{1, 2, 3, 4}, want: false},
		{s: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
	}
	for i := range s {
		require.Equal(t, s[i].want, containsDuplicate(s[i].s), s[i].s)
	}
}

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func containsDuplicate1(nums []int) bool {
	uniq := make(map[int]bool)

	for i := range nums {
		if _, ok := uniq[nums[i]]; ok == true {
			return true
		}
		uniq[nums[i]] = true
	}
	return false
}
