package main

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test49(t *testing.T) {
	type examples struct {
		s    []string
		want [][]string
	}
	s := []examples{
		{
			s:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			s:    []string{"aaa", "bbb", "aaa"},
			want: [][]string{{"aaa", "aaa"}, {"bbb"}},
		},
		{s: []string{""}, want: [][]string{{""}}},
		{s: []string{"a"}, want: [][]string{{"a"}}},
	}
	for i := range s {
		// test might not work due to random map access
		require.Equal(t, s[i].want, groupAnagrams(s[i].s), s[i].s)
	}
}

func groupAnagrams(strs []string) [][]string {
	storeHashes := make(map[string][]int)

	for i, str := range strs {
		hash := makeHash(str)
		storeHashes[hash] = append(storeHashes[hash], i)
	}

	var res [][]string
	for _, indexes := range storeHashes {
		var r []string
		for _, index := range indexes {
			r = append(r, strs[index])
		}
		res = append(res, r)
	}

	return res
}

func makeHash(w string) string {
	// its sorting string, but it's better to count each letter appearance
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
