package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func compareNode(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return compareNode(left.Left, right.Right) && compareNode(left.Right, right.Left)

}

func isSymmetric(root *TreeNode) bool {
	return compareNode(root.Left, root.Right)

}
func main() {
	tree := &TreeNode{Val: 0, Left: &TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}}
	result := isSymmetric(tree)
	fmt.Println(result)

}
