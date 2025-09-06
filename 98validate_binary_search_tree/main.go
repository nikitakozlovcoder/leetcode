package main

import "math"

func main() {

}

/*func isValidBST(root *TreeNode) bool {
	return rec(root, math.MinInt64, math.MaxInt64)
}

func rec(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	return rec(node.Left, min, node.Val) && rec(node.Right, node.Val, max)
}*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	prev := math.MinInt64
	return rec(root, &prev)
}

func rec(node *TreeNode, prev *int) bool {
	if node == nil {
		return true
	}

	left := rec(node.Left, prev)
	if !left || *prev >= node.Val {
		return false
	}

	*prev = node.Val

	right := rec(node.Right, prev)
	return right
}
