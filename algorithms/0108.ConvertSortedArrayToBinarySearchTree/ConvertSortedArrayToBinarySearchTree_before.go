package ConvertSortedArrayToBinarySearchTree

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

func sortedArrayToBST(nums []int) *TreeNode {

	rootNode := &TreeNode{0, nil, nil}

	res := dfs(nums, rootNode)

	return res
}

func dfs(nums []int, currNode *TreeNode) *TreeNode {

	currNodeIdx := len(nums) / 2

	if len(nums) == 0 {
		currNode = nil
		return currNode
	}

	// 末端ノードまで再帰する
	if currNodeIdx != 0 {
		/*
			カレント(ルート)ノードの部分ツリーは・・・
			左部分ツリー：　常にnums[0:k]となる。kはカレントノードとする。
			右部分ツリー：　常にnums[k+1:n]となる。nは、現在の部分ツリーのnums長とする。
		*/
		currNode.Val = nums[currNodeIdx] // 部分ツリーのルートノードへ値セット

		currNode.Left = dfs(nums[0:currNodeIdx], &TreeNode{})
		currNode.Right = dfs(nums[currNodeIdx+1:], &TreeNode{})

	} else {
		// 末端ノードのとき
		currNode = &TreeNode{nums[currNodeIdx], nil, nil}
	}

	return currNode
}
