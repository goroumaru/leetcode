package BinaryTreeLevelOrderTraversal

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

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	stack := make([]*TreeNode, 0)
	depth := make([]int, 0)
	levels := make([][]int, 0)

	// ルートノード追加
	stack = append(stack, root)
	depth = append(depth, 0)

	// 深さ優先探索
	dfs(&stack, &depth, &levels)

	return levels
}

func dfs(stack *[]*TreeNode, depth *[]int, levels *[][]int) {

	currNode := (*stack)[len(*stack)-1]
	*stack = append((*stack)[0:len(*stack)-1], (*stack)[len(*stack):]...)

	currDepth := (*depth)[len(*depth)-1]
	*depth = append((*depth)[0:len(*depth)-1], (*depth)[len(*depth):]...)

	// 出力階層が不足していたら、追加する
	if len(*levels)-1 < currDepth {
		*levels = append(*levels, []int{})
	}
	(*levels)[currDepth] = append((*levels)[currDepth], currNode.Val) // 出力階層へカレントノード追加

	// Right追加
	if currNode.Right != nil {
		*stack = append(*stack, currNode.Right)
		*depth = append(*depth, currDepth+1)
	}

	// Left追加
	if currNode.Left != nil {
		*stack = append(*stack, currNode.Left)
		*depth = append(*depth, currDepth+1)
	}

	if len(*stack) != 0 {
		dfs(stack, depth, levels)
	}
}
