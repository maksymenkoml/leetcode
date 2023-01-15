package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test20(t *testing.T) {
	type examples struct {
		s    string
		want bool
	}
	s := []examples{
		{s: "]", want: false},
		{s: "([)]", want: false},
		{s: "()", want: true},
		{s: "()[]{}", want: true},
		{s: "(]", want: false},
	}
	for i := range s {
		require.Equal(t, s[i].want, isValid2(s[i].s), s[i].s)
	}
}

func isValid2(s string) bool {
	rules := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	var st string
	for _, val := range s {
		if rule, ok := rules[string(val)]; !ok {
			st = st + string(val)
			continue
		} else {
			if st == "" || st[len(st)-1:] != rule {
				return false
			}
		}

		st = st[:len(st)-1]
	}
	return st == ""
}

func isValid1(s string) bool {
	var st string
	for _, val := range s {
		if string(val) == "(" || string(val) == "{" || string(val) == "[" {
			st = st + string(val)
			continue
		}
		if string(val) == ")" {
			if st == "" || st[len(st)-1:] != "(" {
				return false
			}
		}
		if string(val) == "}" {
			if st == "" || st[len(st)-1:] != "{" {
				return false
			}
		}
		if string(val) == "]" {
			if st == "" || st[len(st)-1:] != "[" {
				return false
			}
		}
		st = st[:len(st)-1]
	}
	return st == ""
}
