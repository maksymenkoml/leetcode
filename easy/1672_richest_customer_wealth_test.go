package easy

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1672(t *testing.T) {
	type examples struct {
		accounts [][]int
		want     int
	}
	s := []examples{
		{accounts: [][]int{{1, 5}, {7, 3}, {3, 5}}, want: 10},
		{accounts: [][]int{{1, 2, 3}, {3, 2, 1}}, want: 6},
		{accounts: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, want: 17},
	}
	for _, tst := range s {
		got := maximumWealth(tst.accounts)

		require.Equal(t, tst.want, got)
	}
}

func maximumWealth(accounts [][]int) int {
	wg := sync.WaitGroup{}

	wealths := make(chan int, len(accounts))
	for _, account := range accounts {
		wg.Add(1)
		go func(account []int) {
			defer wg.Done()
			var wealth int
			for _, balance := range account {
				wealth += balance
			}
			wealths <- wealth
		}(account)
	}

	wg.Wait()

	var max int
	for len(wealths) > 0 {
		w := <-wealths
		if w > max {
			max = w
		}
	}

	return max
}

func maximumWealth2(accounts [][]int) int {
	wealths := make(chan int)

	for _, account := range accounts {
		go func(account []int) {
			var wealth int
			for _, balance := range account {
				wealth += balance
			}
			wealths <- wealth
		}(account)
	}

	var max int
	for i := 0; i < len(accounts); i++ {
		w := <-wealths
		if w > max {
			max = w
		}
	}

	return max
}

func maximumWealth1(accounts [][]int) int {
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
