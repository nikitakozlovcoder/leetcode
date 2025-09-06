package main

import (
	"fmt"
)

func main() {
	res := deleteDuplicates2(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
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
		},
	})

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	res := &ListNode{
		Next: head,
	}

	cur := res
	for cur != nil {
		next := cur.Next
		for {
			was := false
			for next != nil && next.Next != nil && next.Val == next.Next.Val {
				was = true
				next = next.Next
			}

			if was {
				next = next.Next
			} else {
				break
			}
		}

		cur.Next = next
		cur = cur.Next
	}

	return res.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
