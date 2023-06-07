package sword_offer

/**
* @Author XuPEngHao
* @Date 2023/6/7 08:55
 */

/*
剑指 Offer 07. 重建二叉树
输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。

假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
示例 1:

Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
示例 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]
限制：

0 <= 节点个数 <= 5000
*/

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

// 先序：根左右
// 中序：左根右
// 建立根节点，划分左右子树
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 中序可以划分 左子树 和 有子树
	inorderMap := make(map[int]int)
	for index, num := range inorder {
		inorderMap[num] = index
	}
	var recurFN func(root, left, right int) *TreeNode

	/*
		preorder := []int{3, 9, 20, 15, 7}
		inorder := []int{9, 3, 15, 20, 7}
	*/
	recurFN = func(root, left, right int) *TreeNode {
		for left > right {
			return nil
		}
		node := &TreeNode{ // 根节点
			Val: preorder[root],
		}

		// 先序节点在中序节的下标 可讲中序 划分为 左子树 和右子树
		i := inorderMap[preorder[root]]
		node.Left = recurFN(root+1, left, i-1)          // 左子树
		node.Right = recurFN(i-left+root+1, i+1, right) // 右子树
		return node
	}

	return recurFN(0, 0, len(inorder)-1)
}
