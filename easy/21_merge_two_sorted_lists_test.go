package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test21(t *testing.T) {
	type examples struct {
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}
	s := []examples{
		{
			list1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			list2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tst := range s {
		require.Equal(t, tst.want, mergeTwoLists(tst.list1, tst.list2))
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	tail := res

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil {
		tail.Next = list2
	}

	return res.Next
}
