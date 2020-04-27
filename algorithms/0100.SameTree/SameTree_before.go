package SameTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	var status bool = true

	queuePtree := make([]*TreeNode, 0, 10000)
	queueQtree := make([]*TreeNode, 0, 10000)

	queuePtree = append(queuePtree, p)
	queueQtree = append(queueQtree, q)

	check(&queuePtree, &queueQtree, &status)

	return status
}

func check(queuePtree *[]*TreeNode, queueQtree *[]*TreeNode, status *bool) {

	// DBSしてカレントノード比較
	if (len(*queuePtree) != 0) && (len(*queueQtree) != 0) {

		currPnode := dbs(queuePtree)
		currQnode := dbs(queueQtree)

		if (currPnode != nil) && (currQnode != nil) {
			if (*currPnode).Val != (*currQnode).Val {
				*status = false
			} else {
				check(queuePtree, queueQtree, status)
			}
		} else if (currPnode == nil) && (currQnode == nil) {
			check(queuePtree, queueQtree, status)
		} else {
			*status = false
		}

	} else if (len(*queuePtree) == 0) && (len(*queueQtree) == 0) {
		*status = true
	} else {
		*status = false
	}
}

func dbs(queue *[]*TreeNode) *TreeNode {

	currNode := (*queue)[0]
	*queue = append((*queue)[:0], (*queue)[1:]...)

	if currNode == nil {
		return nil
	}

	// Add Left
	if currNode.Left != nil {
		*queue = append(*queue, currNode.Left)
	} else {
		*queue = append(*queue, nil)
	}

	// Add Right
	if currNode.Right != nil {
		*queue = append(*queue, currNode.Right)
	} else {
		*queue = append(*queue, nil)
	}

	return currNode
}
*/
