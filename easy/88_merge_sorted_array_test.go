package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test88(t *testing.T) {
	type examples struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}
	s := []examples{
		{nums1: []int{0}, m: 0, nums2: []int{1}, n: 1, want: []int{1}},
		{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, want: []int{1, 2, 2, 3, 5, 6}},
		{nums1: []int{1}, m: 1, nums2: []int{}, n: 0, want: []int{1}},
	}
	for _, tst := range s {
		merge(tst.nums1, tst.m, tst.nums2, tst.n)
		require.Equal(t, tst.want, tst.nums1, tst.want)
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	n--
	m--
	for i := len(nums1) - 1; i >= 0; i-- {
		if nums1[i] != 0 {
			return
		}
		if n >= 0 {
			if m >= 0 {
				if nums2[n] > nums1[m] {
					nums1[i] = nums2[n]
					n--
					continue
				}
			} else {
				nums1[i] = nums2[n]
				n--
				continue
			}
		}
		if m >= 0 {
			nums1[i], nums1[m] = nums1[m], nums1[i]
			m--
		}
	}
}
