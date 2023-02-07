package medium

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test11(t *testing.T) {
	type examples struct {
		height []int
		want   int
	}
	s := []examples{
		{height: []int{2, 3, 4, 5, 18, 17, 6}, want: 17},
		{height: []int{2, 3, 10, 5, 7, 8, 9}, want: 36},
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{height: []int{1, 1}, want: 1},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, maxArea(tst.height), tst.height)
	}
}

func maxArea(height []int) int {
	maxVolume := 0

	l, r := 0, len(height)-1
	for r > l {
		volume := calcVolume(height[l], height[r], r-l)
		if volume > maxVolume {
			maxVolume = volume
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxVolume
}

func maxArea1(height []int) int {
	maxVolume := 0

	for i := range height {
		for j := i + 1; j < len(height); j++ {
			volume := calcVolume(height[i], height[j], j-i)
			if volume > maxVolume {
				maxVolume = volume
			}
		}
	}

	return maxVolume
}

func calcVolume(x, y int, dist int) int {
	return int(math.Min(float64(x), float64(y))) * dist
}
