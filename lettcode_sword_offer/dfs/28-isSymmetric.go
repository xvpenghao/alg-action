package dfs

// @Author XuPEngHao
// @DATE 2022/5/23 19:55

/*
剑指 Offer 28. 对称的二叉树
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3



示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false


限制：

0 <= 节点个数 <= 1000
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetricV2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return dfsIsSymmetricV2(root.Left, root.Right)
}

// 拿着左右两个节点比较
func dfsIsSymmetricV2(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val &&
		dfsIsSymmetricV2(left.Left, right.Right) &&
		dfsIsSymmetricV2(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	numList1 := getLevelTraverse(root)
	dfsIsSymmetric(root)
	numList2 := getLevelTraverse(root)

	for i, _ := range numList1 {
		if numList1[i] != numList2[i] {
			return false
		}
	}
	return true
}

func getLevelTraverse(root *TreeNode) []int {
	queue := []*TreeNode{root}
	var numList []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			numList = append(numList, -9999)
			continue
		}
		numList = append(numList, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	return numList
}

func dfsIsSymmetric(node *TreeNode) {
	if node == nil {
		return
	}
	node.Left, node.Right = node.Right, node.Left
	dfsIsSymmetric(node.Left)
	dfsIsSymmetric(node.Right)
}
