package BinaryTreeZigzagLevelOrderTraversal

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

func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	stack := make([]*TreeNode, 0)
	depth := make([]int, 0)
	levels := make([][]int, 0)

	stack = append(stack, root)
	depth = append(depth, 0)

	// DFS探索
	dfs(&stack, &depth, &levels)

	return levels
}

func dfs(stack *[]*TreeNode, depth *[]int, levels *[][]int) {

	currNode := (*stack)[len(*stack)-1]
	*stack = append((*stack)[:len(*stack)-1], (*stack)[len(*stack):]...)

	currDepth := (*depth)[len(*depth)-1]
	*depth = append((*depth)[:len(*depth)-1], (*depth)[len(*depth):]...)

	// 出力階層が不足のとき
	if len(*levels)-1 < currDepth {
		*levels = append(*levels, []int{})
	}

	/*
		深さによって、追加位置を変更する
		深さ偶数：最後尾へ追加
		深さ奇数：先頭へ追加
	*/
	switch {
	case currDepth%2 == 0:
		(*levels)[currDepth] = append((*levels)[currDepth], currNode.Val) // 最後尾へ
	case currDepth%2 != 0:
		(*levels)[currDepth] = append([]int{currNode.Val}, (*levels)[currDepth]...) // 先頭へ
	}

	if currNode.Right != nil {
		*stack = append(*stack, currNode.Right)
		*depth = append(*depth, currDepth+1)
	}

	if currNode.Left != nil {
		*stack = append(*stack, currNode.Left)
		*depth = append(*depth, currDepth+1)
	}

	if len(*stack) != 0 {
		dfs(stack, depth, levels)
	}
}
