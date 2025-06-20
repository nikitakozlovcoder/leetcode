package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	res := addTwoNumbers(l1, l2)

	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	t := 0
	res := ListNode{Val: 0, Next: nil}
	resTravel := &res
	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		r := (a + b + t) % 10
		t = (a + b + t) / 10
		resTravel.Next = &ListNode{Val: r, Next: nil}
		resTravel = resTravel.Next
	}

	if t != 0 {
		resTravel.Next = &ListNode{Val: t, Next: nil}
	}

	return res.Next
}
