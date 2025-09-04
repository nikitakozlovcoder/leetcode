package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return traverse(p, q)
}

func traverse(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	left := traverse(p.Left, q.Left)
	right := traverse(p.Right, q.Right)

	return left && right
}
