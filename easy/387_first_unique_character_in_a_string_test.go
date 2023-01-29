package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test387(t *testing.T) {
	type examples struct {
		s    string
		want int
	}
	s := []examples{
		{s: "dddccdbba", want: 8},
		{s: "aabb", want: -1},
		{s: "loveleetcode", want: 2},
		{s: "leetcode", want: 0},
	}
	for _, tst := range s {
		got := firstUniqChar(tst.s)

		require.Equal(t, tst.want, got)
	}
}

func firstUniqChar(s string) int {
	type p struct {
		pos int
		run rune
	}
	store := make(map[rune]int, 26)
	var pos []p

	for i, r := range s {
		if _, ok := store[r]; !ok {
			pos = append(pos, p{i, r})
		}
		store[r]++
	}
	for i := range pos {
		if store[pos[i].run] == 1 {
			return pos[i].pos
		}
	}
	return -1
}
