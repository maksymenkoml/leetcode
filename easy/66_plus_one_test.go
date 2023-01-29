package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test66(t *testing.T) {
	type examples struct {
		digits []int
		want   []int
	}
	s := []examples{
		{digits: []int{8, 9, 9, 9}, want: []int{9, 0, 0, 0}},
		{digits: []int{1, 2, 3}, want: []int{1, 2, 4}},
		{digits: []int{4, 3, 2, 1}, want: []int{4, 3, 2, 2}},
		{digits: []int{9}, want: []int{1, 0}},
	}
	for _, tst := range s {
		got := plusOne(tst.digits)

		require.Equal(t, tst.want, got)
	}
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}

func plusOne1(digits []int) []int {
	i := len(digits) - 1
	var isTen bool
	for {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
			isTen = true
		}

		if i == 0 {
			if isTen {
				digits = append([]int{1}, digits...)
			}
			return digits
		}

		i--
	}
}
