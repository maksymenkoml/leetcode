package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test26(t *testing.T) {
	type examples struct {
		nums     []int
		wantK    int
		wantNums []int
	}
	s := []examples{
		{nums: []int{0, 1, 2, 2, 2, 2, 2}, wantK: 3, wantNums: []int{0, 1, 2}},
		{nums: []int{1}, wantK: 1, wantNums: []int{1}},
		{nums: []int{1, 1, 2}, wantK: 2, wantNums: []int{1, 2}},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, wantK: 5, wantNums: []int{0, 1, 2, 3, 4}},
	}
	for _, tst := range s {
		got := removeDuplicates(tst.nums)

		require.Equal(t, tst.wantK, got)
		require.Equal(t, tst.wantNums, tst.nums[:got])
	}
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

func removeDuplicates2(nums []int) int {
	store := make(map[int]bool, len(nums)+1)
	ln := len(nums)
	last := nums[len(nums)-1]
	store[nums[0]] = true
	i := 1
	for {
		if i >= ln {
			break
		}
		if nums[i-1] > nums[i] {
			break
		}
		if _, ok := store[nums[i]]; !ok {
			store[nums[i]] = true
			i++
			continue
		}
		if nums[i] == last {
			break
		}
		nums = append(nums[:i], append(nums[i+1:], nums[i])...)
	}
	return i
}
