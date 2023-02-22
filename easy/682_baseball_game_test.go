package easy

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test682(t *testing.T) {
	type examples struct {
		operations []string
		want       int
	}
	s := []examples{
		{operations: []string{"5", "2", "C", "D", "+"}, want: 30},
		{operations: []string{"5", "-2", "4", "C", "D", "9", "+", "+"}, want: 27},
		{operations: []string{"1", "C"}, want: 0},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, calPoints(tst.operations), tst.operations)
	}
}

func calPoints(operations []string) int {
	var sum int
	var store []int
	for _, op := range operations {
		i, err := strconv.Atoi(op)
		if err != nil {
			switch op {
			case "C":
				last := store[len(store)-1]
				sum -= last
				store = store[:len(store)-1]

			case "D":
				last := store[len(store)-1]
				res := last * 2
				sum += res
				store = append(store, res)

			case "+":
				last := store[len(store)-1]
				last2 := store[len(store)-2]
				res := last + last2
				sum += res
				store = append(store, res)

			}
			continue
		}
		sum += i
		store = append(store, i)
	}

	return sum
}
