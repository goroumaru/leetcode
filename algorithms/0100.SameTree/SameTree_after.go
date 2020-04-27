package SameTree

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

func isSameTree(p *TreeNode, q *TreeNode) bool {

	// nilチェック
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}

	// 同値チェック
	if p.Val != q.Val {
		return false
	}

	/*
		同値でなかったら、次のノードへ遷移
		そのとき、左と右ノードが"ともに"一致しているかチェックする
	*/
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
