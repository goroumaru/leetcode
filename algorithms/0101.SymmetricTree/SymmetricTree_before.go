package SymmetricTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return checkSymmetric(root, root)
}

func checkSymmetric(nodeA *TreeNode, nodeB *TreeNode) bool {

	// 末端ノードに到達したとき、両ノードとも先がなければTrueとする
	if nodeA == nil && nodeB == nil {
		return true
	}

	/*
		以下、Falseとなる条件
		同値となっていても、末端ノードとなるまではFalseのまま
	*/
	if nodeA == nil && nodeB != nil {
		return false
	}
	if nodeA != nil && nodeB == nil {
		return false
	}
	if nodeA.Val != nodeB.Val {
		return false
	}

	return checkSymmetric(nodeA.Left, nodeB.Right) && checkSymmetric(nodeA.Right, nodeB.Left)
}
