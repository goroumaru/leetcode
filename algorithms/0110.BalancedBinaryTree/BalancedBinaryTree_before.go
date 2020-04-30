package BalancedBinaryTree

import (
	"math"
)

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

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	queue := make([]*TreeNode, 0)
	depth := make([]int, 0)

	queue = append(queue, root)
	depth = append(depth, 0)

	ok, _ := dfs(&queue, &depth)

	return ok
}

func dfs(queue *[]*TreeNode, depth *[]int) (bool, int) {

	if len(*queue) == 0 {
		return true, 0
	}

	currNode := (*queue)[len(*queue)-1]
	*queue = append((*queue)[:len(*queue)-1], (*queue)[len(*queue):]...)

	currDepth := (*depth)[len(*depth)-1]
	*depth = append((*depth)[:len(*depth)-1], (*depth)[len(*depth):]...)

	var leftOk bool = true
	leftDepth := currDepth
	if currNode.Left != nil {
		*queue = append(*queue, currNode.Left)
		*depth = append(*depth, currDepth+1)
		leftOk, leftDepth = dfs(queue, depth)
	}

	var rightOk bool = true
	rightDepth := currDepth
	if currNode.Right != nil {
		*queue = append(*queue, currNode.Right)
		*depth = append(*depth, currDepth+1)
		rightOk, rightDepth = dfs(queue, depth)
	}

	/*
		左右ノードバランスが取れているか
	*/
	var balancedOk bool = true
	if !(rightOk) || !(leftOk) {
		balancedOk = false
	}
	if math.Abs(float64(rightDepth-leftDepth)) > 1 {
		balancedOk = false
	}

	/*
		Left , Rightのうち、末端ノードまでが長い方を選択する
	*/
	currMaxDepth := leftDepth
	if rightDepth > leftDepth {
		currMaxDepth = rightDepth
	}

	// 末端ノードのとき
	if currNode.Left == nil && currNode.Right == nil {
		currMaxDepth = currDepth
	}

	return balancedOk, currMaxDepth
}
