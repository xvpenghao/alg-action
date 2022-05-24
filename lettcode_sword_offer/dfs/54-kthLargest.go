package dfs

import (
	"sort"
)

// @Author XuPEngHao
// @DATE 2022/5/24 08:53

/*
剑指 Offer 54. 二叉搜索树的第k大节点
给定一棵二叉搜索树，请找出其中第 k 大的节点的值。



示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4


限制：

1 ≤ k ≤ 二叉搜索树元素个数
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var k, val int
// 充分的学会利用已知条件
func kthLargestV3(root *TreeNode, k2 int) int {
	if root == nil {
		return 0
	}
	k = k2
	midTraverseV3(root)
	return val
}

func midTraverseV3(node *TreeNode) {
	if node == nil {
		return
	}
	midTraverseV3(node.Right)
	if k == 0 {
		return
	}
	k--
	val = node.Val
	midTraverseV3(node.Left)
}

func kthLargestV2(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	// 1 ≤ k ≤ 二叉搜索树元素个数 根据题意，这里就不考虑，下标越界可能性了。
	numList := midTraverse(root)
	return numList[len(numList)-k]
}

// 二叉树搜素数中序遍历，是排好序的
func midTraverse(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	var numList []int
	numList = append(numList, midTraverse(node.Left)...)
	numList = append(numList, node.Val)
	numList = append(numList, midTraverse(node.Right)...)
	return numList
}

func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	// 1 ≤ k ≤ 二叉搜索树元素个数 根据题意，这里就不考虑，下标越界可能性了。
	numList := levelTraverse(root)
	sort.SliceStable(numList, func(i, j int) bool {
		return numList[i] > numList[j]
	})
	return numList[k-1]
}

func levelTraverse(root *TreeNode) []int {
	var numList []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		numList = append(numList, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return numList
}
