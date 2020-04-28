package MaximumDepthOfBinaryTree

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

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	stack := make([]*TreeNode, 0)
	depth := make([]int, 0)

	stack = append(stack, root)
	depth = append(depth, 0)

	var maxDepth int
	dfs(&stack, &depth, &maxDepth)

	return maxDepth + 1
}

func dfs(stack *[]*TreeNode, depth *[]int, maxDepth *int) {

	currNode := (*stack)[len(*stack)-1]
	*stack = append((*stack)[:len(*stack)-1], (*stack)[len(*stack):]...)

	currDepth := (*depth)[len(*depth)-1]
	*depth = append((*depth)[:len(*depth)-1], (*depth)[len(*depth):]...)

	// 最大深さ判定
	if currDepth > *maxDepth {
		*maxDepth = currDepth
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
		dfs(stack, depth, maxDepth)
	}

}
