package main

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test242(t *testing.T) {
	type examples struct {
		s    string
		t    string
		want bool
	}
	s := []examples{
		{s: "anagram", t: "nagaram", want: true},
		{s: "rat", t: "cat", want: false},
	}
	for i := range s {
		require.Equal(t, s[i].want, isAnagram(s[i].s, s[i].t))
	}
}

func isAnagram(s string, t string) bool {
	sA := strings.Split(s, "")
	tA := strings.Split(t, "")

	sort.Strings(sA)
	sort.Strings(tA)

	sB := strings.Join(sA, "")
	tB := strings.Join(tA, "")

	return sB == tB
}
