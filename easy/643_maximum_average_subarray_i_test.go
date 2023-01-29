package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test643(t *testing.T) {
	type examples struct {
		nums []int
		k    int
		want float64
	}
	s := []examples{
		{nums: []int{0, 1, 1, 3, 3}, k: 4, want: 2.00000},
		{nums: []int{1, 12, -5, -6, 50, 3}, k: 4, want: 12.75000},
		{nums: []int{5}, k: 1, want: 5},
	}
	for _, tst := range s {
		got := findMaxAverage(tst.nums, tst.k)

		require.Equal(t, tst.want, got, tst.nums)
	}
}

func findMaxAverage(nums []int, k int) float64 {
	var sum int
	for j := 0; j < k; j++ {
		sum += nums[j]
	}
	maxSum := sum

	for i := 1; i <= len(nums)-k; i++ {
		sum -= nums[i-1]
		sum += nums[i+k-1]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}
