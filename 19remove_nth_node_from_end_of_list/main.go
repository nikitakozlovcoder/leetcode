package main

import "fmt"

func main() {
	list := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 2},
	}

	result := removeNthFromEnd(list, 2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head
	for range n {
		fast = fast.Next
	}

	if fast == nil {
		return slow.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}
