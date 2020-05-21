package leetcode

/*
Problem statement

In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.
Two nodes of a binary tree are cousins if they have the same depth, but have different parents.

We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.
Return true if and only if the nodes corresponding to the values x and y are cousins.

Input: root = [1,2,3,4], x = 4, y = 3
Output: false

Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
Output: true
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	parentX *TreeNode
	parentY *TreeNode
	depthX  int
	depthY  int
)

func isCousins(root *TreeNode, x int, y int) bool {

	getParentAndDepth(root, nil, x, y, 0)
	if depthX == depthY && parentX != parentY {
		return true
	}
	return false
}

func getParentAndDepth(root *TreeNode, parent *TreeNode, x int, y int, depth int) {
	if root == nil {
		return
	}
	if root.Val == x {
		parentX = parent
		depthX = depth
	} else if root.Val == y {
		parentY = parent
		depthY = depth
	} else {
		getParentAndDepth(root.Left, root, x, y, depth+1)
		getParentAndDepth(root.Right, root, x, y, depth+1)

	}
}
