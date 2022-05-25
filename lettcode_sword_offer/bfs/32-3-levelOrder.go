package bfs

import "sort"

// @Author XuPEngHao
// @DATE 2022/5/25 12:03

/*
剑指 Offer 32 - III. 从上到下打印二叉树 III
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。



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
  [20,9],
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

// 层次遍历 + 双端队列【奇-左，偶-右】
// 层次遍历+ 奇偶判定+排序
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var numList [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var tempNumList []int
		var tempNodeList []*TreeNode
		for _, node := range queue {
			tempNumList = append(tempNumList, node.Val)
			if node.Left != nil {
				tempNodeList = append(tempNodeList, node.Left)
			}
			if node.Right != nil {
				tempNodeList = append(tempNodeList, node.Right)
			}
		}
		if len(numList)%2 == 1 {
			sort.SliceStable(tempNumList, func(i, j int) bool {
				return i > j
			})
		}
		numList = append(numList, tempNumList)
		queue = tempNodeList
	}
	return numList
}
