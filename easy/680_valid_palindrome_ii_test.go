package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test128(t *testing.T) {
	type examples struct {
		s    string
		want bool
	}
	s := []examples{
		{s: "abc", want: false},
		{s: "ebcbbececabbacecbbcbe", want: true},
		{s: "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga", want: true},
		{s: "cbbcc", want: true},
		{s: "eedede", want: true},
		{s: "eeccccbebaeeabebccceea", want: false},
		{s: "aba", want: true},
		{s: "abca", want: true},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, validPalindrome(tst.s), tst.s)
	}
}

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	once := false
	twice := false

NextPass:
	for r > l {
		if s[l] != s[r] {
			if once {
				twice = true
				break NextPass
			}
			once = true
			if s[l+1] != s[r] {
				if s[l] != s[r-1] {
					return false
				} else {
					r--
				}
			} else {
				l++
			}
		}
		l++
		r--
	}

	if !twice {
		return true
	}

	l, r = 0, len(s)-1
	once = false
	for r > l {
		if s[l] != s[r] {
			if once {
				return false
			}
			once = true
			if s[l] != s[r-1] {
				if s[l+1] != s[r] {
					return false
				} else {
					l++
				}
			} else {
				r--
			}
		}
		l++
		r--
	}

	return true
}
