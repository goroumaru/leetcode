package BinaryTreeLevelOrderTraversal2

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

func levelOrderBottom(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	queue := make([]*TreeNode, 0)
	depth := make([]int, 0)
	levels := make([][]int, 0)

	queue = append(queue, root)
	depth = append(depth, 0)
	levels = append(levels, []int{})

	// 幅優先探索
	bfs(&queue, &depth, &levels)

	return levels
}

func bfs(queue *[]*TreeNode, depth *[]int, levels *[][]int) {

	currNode := (*queue)[0]
	*queue = append((*queue)[1:len(*queue)], (*queue)[len(*queue):]...)

	currDepth := (*depth)[0]
	*depth = append((*depth)[1:len(*depth)], (*depth)[len(*depth):]...)

	// 出力レベルの階層不足していたら、追加する
	if len(*levels)-1 < currDepth {
		*levels = append([][]int{[]int{}}, *levels...) // 先頭へ追加する(空配列と結合する)
	}

	// 深い階層ほど先頭へ追加する
	(*levels)[0] = append((*levels)[0], currNode.Val) // 該当する階層へカレントノード追加

	if currNode.Left != nil {
		*queue = append(*queue, currNode.Left)
		*depth = append(*depth, currDepth+1)
	}

	if currNode.Right != nil {
		*queue = append(*queue, currNode.Right)
		*depth = append(*depth, currDepth+1)
	}

	if len(*queue) != 0 {
		bfs(queue, depth, levels)
	}
}
