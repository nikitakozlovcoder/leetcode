package main

import (
	"fmt"
	"math"
)

func main() {
	val := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 2,
			},
		},
	}

	recoverTree(val)
	fmt.Println(val)
}

func recoverTree(root *TreeNode) {
	rec(root, math.MinInt64, math.MaxInt64)
}

func rec(node *TreeNode, min, max int) (*TreeNode, int) {
	if node == nil {
		return nil, 0
	}

	if node.Val <= min {
		return node, min
	}

	if node.Val >= max {
		return node, max
	}

	res, t := rec(node.Left, min, node.Val)
	if res != nil {
		if node.Val == t {
			node.Val, res.Val = res.Val, node.Val
			return nil, 0
		}

		return res, t
	}

	res, t = rec(node.Right, node.Val, max)
	if res != nil {
		if node.Val == t {
			node.Val, res.Val = res.Val, node.Val
			return nil, 0
		}

		return res, t
	}

	return nil, 0
}*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
