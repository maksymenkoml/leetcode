package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test344(t *testing.T) {
	type examples struct {
		s    []byte
		want []byte
	}
	s := []examples{
		{s: []byte{'h', 'e', 'l', 'l', 'o'}, want: []byte{'o', 'l', 'l', 'e', 'h'}},
		{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}, want: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for _, tst := range s {
		reverseString(tst.s)
		require.Equal(t, tst.want, tst.s, tst.s)
	}
}

func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
