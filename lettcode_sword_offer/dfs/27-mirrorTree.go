package dfs

// @Author XuPEngHao
// @DATE 2022/5/22 16:20

/*
剑指 Offer 27. 二叉树的镜像
请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

示例 1：

输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTreeV2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := root

	dfsMirrorTree(node)
	return root
}

func dfsMirrorTree(node *TreeNode) {
	if node == nil {
		return
	}
	// 交换节点
	node.Left, node.Right = node.Right, node.Left //  上层已经交换了，一
	dfsMirrorTree(node.Left)                      // 可以叫dfs，一直去探索，寻找可以走的路
	dfsMirrorTree(node.Right)
}

// 借助队列，完成二叉树的交换
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		// 交换
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}
