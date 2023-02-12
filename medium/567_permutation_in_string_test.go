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
		{s1: "abc", s2: "ccccbbbbaaaa", want: false},
		{s1: "ab", s2: "eidbaooo", want: true},
		{s1: "ab", s2: "eidboaoo", want: false},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, checkInclusion(tst.s1, tst.s2), tst.s1+" : "+tst.s2)
	}
}

func checkInclusion(s1 string, s2 string) bool {
	val := make(map[byte]int, 26)
	for i := range s1 {
		val[s1[i]]++
	}

	l := 0
	calcVal := make(map[byte]int, 26)
	for r := 0; r < len(s2); r++ {
		calcVal[s2[r]]++

		if r-l > len(s1)-1 {
			calcVal[s2[l]]--
			l++
		}

		if cmpMaps(val, calcVal) {
			return true
		}
	}

	return false
}

func cmpMaps(m1, m2 map[byte]int) bool {
	for i := range m1 {
		if m1[i] != m2[i] {
			return false
		}
	}
	return true
}
