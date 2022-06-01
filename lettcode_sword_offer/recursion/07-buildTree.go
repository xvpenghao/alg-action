package recursion

// @Author XuPEngHao
// @DATE 2022/5/28 16:26

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

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 { // 判断先序和中序都一样
		return nil
	}
	// 根据先序和中序的结合，找出，划分，左右子树
	root := &TreeNode{Val: preorder[0], Left: nil, Right: nil}
	i := 0
	for ; i < len(inorder); i++ { //  从中序里面找出，先序的根节点
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 我们知道根了，就能清除的，划分，根的左右边的子树有那些了，
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	// 排查左边的就是右边的
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
