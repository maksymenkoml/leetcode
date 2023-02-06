package easy

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1984(t *testing.T) {
	type examples struct {
		nums []int
		k    int
		want int
	}
	s := []examples{
		{nums: []int{41900, 69441, 94407, 37498, 20299, 10856, 36221, 2231, 54526, 79072, 84309, 76765, 92282, 13401, 44698, 17586, 98455, 47895, 98889, 65298, 32271, 23801, 83153, 12186, 7453, 79460, 67209, 54576, 87785, 47738, 40750, 31265, 77990, 93502, 50364, 75098, 11712, 80013, 24193, 35209, 56300, 85735, 3590, 24858, 6780, 50086, 87549, 7413, 90444, 12284, 44970, 39274, 81201, 43353, 75808, 14508, 17389, 10313, 90055, 43102, 18659, 20802, 70315, 48843, 12273, 78876, 36638, 17051, 20478}, k: 5, want: 1428},
		{nums: []int{87063, 61094, 44530, 21297, 95857, 93551, 9918}, k: 6, want: 74560},
		{nums: []int{9}, k: 1, want: 0},
		{nums: []int{9, 4, 1, 7}, k: 2, want: 2},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, minimumDifference(tst.nums, tst.k), tst.nums)
	}
}

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)

	l, r := 0, k-1
	min := nums[r] - nums[l]
	for i := k; i < len(nums); i++ {
		l++
		r++
		newMin := nums[r] - nums[l]
		if newMin < min {
			min = newMin
		}
	}

	return min
}
