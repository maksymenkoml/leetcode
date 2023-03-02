package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test567(t *testing.T) {
	type examples struct {
		s1   string
		s2   string
		want bool
	}
	s := []examples{
		{s1: "abcdxabcde", s2: "abcdeabcdx", want: true},
		{s1: "ab", s2: "eidbaooo", want: true},
		{s1: "abc", s2: "ccccbbbbaaaa", want: false},
		{s1: "ab", s2: "eidboaoo", want: false},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, checkInclusion(tst.s1, tst.s2), tst.s1+" : "+tst.s2)
	}
}

func checkInclusion(s1 string, s2 string) bool {
	val := make(map[byte]int, len(s1))
	for i := range s1 {
		val[s1[i]]++
	}

	match := 0
	l := 0
	calcVal := make(map[byte]int, 'z'-'a')
	for r := 0; r < len(s2); r++ {
		calcVal[s2[r]]++

		if r-l > len(s1)-1 {
			if val[s2[l]] == calcVal[s2[l]] {
				match--
			}
			calcVal[s2[l]]--
			l++
		}

		if val[s2[r]] == calcVal[s2[r]] {
			match++
		}

		if match == len(val) {
			return true
		}
	}

	return false
}
