package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test206(t *testing.T) {
	type examples struct {
		head *ListNode
		want *ListNode
	}
	s := []examples{
		{
			head: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  -2,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: -2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val:  0,
							Next: nil,
						},
					},
				},
			},
		},
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
		{
			head: nil,
			want: nil,
		},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, reverseList(tst.head))
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for {
		if curr == nil {
			break
		}

		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return prev
}
