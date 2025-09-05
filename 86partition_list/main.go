package main

import (
	"fmt"
)

func main() {
	res := partition(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
		},
	}, 3)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	lessNode := &ListNode{}
	greaterNode := &ListNode{}
	lessCur := lessNode
	greaterCur := greaterNode

	cur := head
	for cur != nil {
		if cur.Val < x {
			lessCur.Next = cur
			lessCur = lessCur.Next
		} else {
			greaterCur.Next = cur
			greaterCur = greaterCur.Next
		}

		cur = cur.Next
	}

	lessCur.Next = greaterNode.Next
	greaterCur.Next = nil
	return lessNode.Next
}
