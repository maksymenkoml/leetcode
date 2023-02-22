package medium

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test150(t *testing.T) {
	type examples struct {
		tokens []string
		want   int
	}
	s := []examples{
		{tokens: []string{"2", "1", "+", "3", "*"}, want: 9},
		{tokens: []string{"4", "13", "5", "/", "+"}, want: 6},
		{tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, want: 22},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, evalRPN(tst.tokens), tst.tokens)
	}
}

func evalRPN(tokens []string) int {
	var store []int
	for _, token := range tokens {
		i, err := strconv.Atoi(token)
		if err != nil {
			// pop
			last2 := store[len(store)-1]
			store = store[:len(store)-1]
			// pop
			last1 := store[len(store)-1]
			store = store[:len(store)-1]

			switch token {
			case "+":
				store = append(store, last1+last2)
			case "-":
				store = append(store, last1-last2)
			case "*":
				store = append(store, last1*last2)
			case "/":
				store = append(store, last1/last2)
			}
			continue
		}
		// push
		store = append(store, i)
	}

	return store[len(store)-1]
}
