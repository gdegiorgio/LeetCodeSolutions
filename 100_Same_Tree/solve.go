package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if (p.Val != q.Val) ||
		(p.Right == nil && q.Right != nil) ||
		(p.Right != nil && q.Right == nil) ||
		(p.Left == nil && q.Left != nil) ||
		(p.Left != nil && q.Left == nil) {
		return false
	} else {
		// Means that nodes are leaves, otherwise false is returned
		if p.Right == nil {
			return true
		} else {
			return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
		}
	}
}
