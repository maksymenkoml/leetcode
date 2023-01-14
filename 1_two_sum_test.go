package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	type examples struct {
		nums   []int
		target int
		want   []int
	}
	s := []examples{
		{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	}
	for i := range s {
		want := s[i].want
		sort.Ints(want)

		got := twoSum3(s[i].nums, s[i].target)
		sort.Ints(got)

		require.Equal(t, want, got)
	}
}

func twoSum3(nums []int, target int) []int {
	var positions = make(map[int]int)
	for cur, num := range nums {
		if pos, ok := positions[target-num]; ok {
			return []int{cur, pos}
		}
		positions[num] = cur
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	var pos = make(map[int]int)
	for i := range nums {
		pos[nums[i]] = i
	}
	for k := range nums {
		lft := target - nums[k]

		if p, ok := pos[lft]; ok && p != k {
			return []int{k, p}
		}
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[j]+nums[i] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
