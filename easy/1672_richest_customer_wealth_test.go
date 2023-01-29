package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1672(t *testing.T) {
	type examples struct {
		accounts [][]int
		want     int
	}
	s := []examples{
		{accounts: [][]int{{1, 2, 3}, {3, 2, 1}}, want: 6},
		{accounts: [][]int{{1, 5}, {7, 3}, {3, 5}}, want: 10},
		{accounts: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, want: 17},
	}
	for _, tst := range s {
		got := maximumWealth(tst.accounts)

		require.Equal(t, tst.want, got)
	}
}

func maximumWealth(accounts [][]int) int {
	customersWealth := make([]int, len(accounts))

	var max int
	for i, account := range accounts {
		for _, balance := range account {
			customersWealth[i] += balance
		}
		if customersWealth[i] > max {
			max = customersWealth[i]
		}
	}

	return max
}
