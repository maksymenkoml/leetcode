package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test238(t *testing.T) {
	type examples struct {
		nums []int
		want []int
	}
	s := []examples{
		{nums: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
		{nums: []int{0, 0}, want: []int{0, 0}},
		{nums: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0}},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, productExceptSelf(tst.nums), tst.nums)
	}
}

func productExceptSelf(nums []int) []int {
	products := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		products[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		products[i] *= postfix
		postfix *= nums[i]
	}

	return products
}

func productExceptSelf1(nums []int) []int {
	zero := false
	twoZero := false
	total := 1
	for _, num := range nums {
		if num == 0 {
			if zero {
				twoZero = true
				break
			}
			zero = true
			continue
		}
		total *= num
	}

	res := make([]int, len(nums))

	if twoZero {
		return res
	}

	for i, num := range nums {
		if num == 0 {
			res[i] = total
			continue
		}
		if zero {
			res[i] = 0
			continue
		}
		res[i] = total / num
	}
	return res
}
