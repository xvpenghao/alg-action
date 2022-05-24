package dfs

import "math"

// @Author XuPEngHao
// @DATE 2022/5/24 14:25

/*
剑指 Offer 55 - I. 二叉树的深度
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

例如：

给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。



提示：

节点总数 <= 10000
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 后序遍历，然后比较左右，+1
	return int(math.Max(float64(maxDepth2(root.Left)), float64(maxDepth2(root.Right)))) + 1
}

// 如何区分这一层呢？
// 将下一层的节点存好，上一层的节点遍历完毕后，
// 下层覆盖上层,把层次的节点存储起来
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var tmp []*TreeNode
		for _, node := range queue { // 将这一层。放到临时
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		depth++
		queue = tmp
	}
	return depth
}
