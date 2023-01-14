package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test121(t *testing.T) {
	type examples struct {
		s    []int
		want int
	}
	s := []examples{
		{s: []int{7, 1, 5, 3, 6, 4}, want: 5},
		{s: []int{7, 6, 4, 3, 1}, want: 0},
		{s: []int{1, 2}, want: 1},
	}
	for i := range s {
		require.Equal(t, s[i].want, maxProfit(s[i].s), s[i].s)
	}
}

func maxProfit(prices []int) int {
	profitMax := 0
	l, r := 0, 1
	for r < len(prices) {
		if prices[l] < prices[r] {
			if (prices[r] - prices[l]) > profitMax {
				profitMax = prices[r] - prices[l]
			}
		} else {
			l = r
		}
		r++
	}
	return profitMax
}
