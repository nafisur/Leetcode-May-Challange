package leetcode

/*
Construct Binary Search Tree from Preorder Traversal
Solution
Return the root node of a binary search tree that matches the given preorder traversal.

(Recall that a binary search tree is a binary tree where for every node, any descendant of node.left has a value < node.val, and any descendant of node.right has a value > node.val.  Also recall that a preorder traversal displays the value of the node first, then traverses node.left, then traverses node.right.)

It's guaranteed that for the given test cases there is always possible to find a binary search tree with the given requirements.

Input: [8,5,1,7,10,12]
Output: [8,5,10,1,7,null,12]
*/

/*
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := preorder[0]
	index := len(preorder)

	for i := range preorder {
		if root < preorder[i] {
			index = i
			break
		}
	}
	return &TreeNode{root, bstFromPreorder(preorder[1:index]), bstFromPreorder(preorder[index:])}
}
