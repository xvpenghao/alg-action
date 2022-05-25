package bfs

// @Author XuPEngHao
// @DATE 2022/5/25 11:51

/*
剑指 Offer 32 - II. 从上到下打印二叉树 II
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]


提示：

节点总数 <= 1000
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var numList [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var tmpNumList []int
		var tmpNodeList []*TreeNode
		for _, node := range queue { // 将下一层的节点。临时存储起来
			tmpNumList = append(tmpNumList, node.Val)
			if node.Left != nil {
				tmpNodeList = append(tmpNodeList, node.Left)
			}
			if node.Right != nil {
				tmpNodeList = append(tmpNodeList, node.Right)
			}
		}
		numList = append(numList, tmpNumList)
		queue = tmpNodeList
	}
	return numList
}
