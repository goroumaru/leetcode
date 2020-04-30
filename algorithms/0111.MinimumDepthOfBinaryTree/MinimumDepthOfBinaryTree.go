package MinimumDepthOfBinaryTree

import "math"

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

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	stack := make([]*TreeNode, 0)
	depth := make([]int, 0)

	stack = append(stack, root)
	depth = append(depth, 0)

	var minDepth int = math.MaxInt32
	dfs(&stack, &depth, &minDepth)

	return minDepth + 1
}

func dfs(stack *[]*TreeNode, depth *[]int, minDepth *int) {

	if len(*stack) == 0 {
		return
	}

	currNode := (*stack)[len(*stack)-1]
	*stack = append((*stack)[:len(*stack)-1], (*stack)[len(*stack):]...)

	currDepth := (*depth)[len(*depth)-1]
	*depth = append((*depth)[:len(*depth)-1], (*depth)[len(*depth):]...)

	if currNode.Right != nil {
		*stack = append(*stack, currNode.Right)
		*depth = append(*depth, currDepth+1)
		dfs(stack, depth, minDepth)
	}

	if currNode.Left != nil {
		*stack = append(*stack, currNode.Left)
		*depth = append(*depth, currDepth+1)
		dfs(stack, depth, minDepth)
	}

	// 末端（葉）ノードのとき
	if currNode.Left == nil && currNode.Right == nil {
		if currDepth < *minDepth {
			*minDepth = currDepth
		}
	}
}
