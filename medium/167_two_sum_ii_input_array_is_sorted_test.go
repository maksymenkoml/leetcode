package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test167(t *testing.T) {
	type examples struct {
		numbers []int
		target  int
		want    []int
	}
	s := []examples{
		{numbers: []int{5, 25, 75}, target: 100, want: []int{2, 3}},
		{numbers: []int{-1, -1, 1}, target: -2, want: []int{1, 2}},
		{numbers: []int{-1, -1, 1}, target: -2, want: []int{1, 2}},
		{numbers: []int{-4, -3, 3, 4, 90}, target: 0, want: []int{1, 4}},
		{numbers: []int{0, 0, 3, 4}, target: 0, want: []int{1, 2}},
		{numbers: []int{2, 7, 11, 15}, target: 9, want: []int{1, 2}},
		{numbers: []int{2, 3, 4}, target: 6, want: []int{1, 3}},
		{numbers: []int{-1, 0}, target: -1, want: []int{1, 2}},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, twoSum(tst.numbers, tst.target), tst.numbers)
	}
}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for numbers[l]+numbers[r] != target {
		for numbers[l]+numbers[r] > target {
			r--
		}
		if numbers[l]+numbers[r] == target {
			break
		}
		l++
	}

	return []int{l + 1, r + 1}
}

func twoSum1(numbers []int, target int) []int {
	l, r := 0, 1

	for numbers[l]+numbers[r] != target {
		for numbers[l]+numbers[r] < target && r < len(numbers)-1 {
			r++
		}
		if numbers[l]+numbers[r] == target {
			break
		}
		l++
		r = l + 1
	}

	if l > r {
		return []int{r + 1, l + 1}
	}
	return []int{l + 1, r + 1}
}
