package easy

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test412(t *testing.T) {
	type examples struct {
		n    int
		want []string
	}
	s := []examples{
		{n: 3, want: []string{"1", "2", "Fizz"}},
		{n: 5, want: []string{"1", "2", "Fizz", "4", "Buzz"}},
		{n: 15, want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}
	for _, tst := range s {
		got := fizzBuzz(tst.n)

		require.Equal(t, tst.want, got)
	}
}

func fizzBuzz(n int) []string {
	res := make([]string, n+1, n+1)
	var r string
	for i := 1; i <= n; i++ {
		r = ""
		if i%3 == 0 {
			r = r + "Fizz"
		}
		if i%5 == 0 {
			r = r + "Buzz"
		}
		if r == "" {
			r = strconv.Itoa(i)
		}
		res[i] = r
	}
	return res[1:]
}
