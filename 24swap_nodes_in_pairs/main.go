package main

import "fmt"

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}

	head := swapPairs(list)
	for head != nil {
		fmt.Println(head)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	res := &ListNode{Val: 0, Next: head}
	orange := res
	green := head
	for green != nil && green.Next != nil {
		orange.Next = green.Next
		green.Next = orange.Next.Next
		orange.Next.Next = green

		orange = green
		green = green.Next
	}

	return res.Next
}
