package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test125(t *testing.T) {
	type examples struct {
		s    string
		want bool
	}
	s := []examples{
		{s: "a12A", want: false},
		{s: "a11A", want: true},
		{s: "a121A", want: true},
		{s: "A man, a plan, a canal: Panama", want: true},
		{s: "0P", want: false},
		{s: "a.", want: true},
		{s: "race a car", want: false},
		{s: " ", want: true},
	}
	for i := range s {
		require.Equal(t, s[i].want, isPalindrome2(s[i].s), s[i].s)
	}
}

func isPalindrome2(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if !isAlNum(rune(s[l])) {
			l++
			continue
		}
		if !isAlNum(rune(s[r])) {
			r--
			continue
		}

		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}

		l++
		r--
	}

	return true
}

func isPalindrome1(s string) bool {
	s = strings.ToLower(s)

	var sanitizedL string
	var sanitizedR string
	for i := 0; i < len(s); i++ {
		l := s[i]
		r := s[len(s)-i-1]

		if i >= len(s)/2 {
			break
		}

		if isAlNum(rune(l)) {
			sanitizedL = sanitizedL + string(l)
		}

		if isAlNum(rune(r)) {
			sanitizedR = sanitizedR + string(r)
		}
	}

	return sanitizedL == sanitizedR
}

func isAlNum(s rune) bool {
	if int(s) >= 'A' && int(s) <= 'Z' {
		return true
	}
	if int(s) >= 'a' && int(s) <= 'z' {
		return true
	}
	if s >= '0' && s <= '9' {
		return true
	}
	return false
}
