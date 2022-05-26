package tree

import "math"

// @Author XuPEngHao
// @DATE 2022/5/25 13:00

/*
剑指 Offer 68 - I. 二叉搜索树的最近公共祖先
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]





示例 1:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。
示例 2:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。


说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉搜索树中。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 最近祖先节点，（最深）
	for root != nil {
		if root.Val < p.Val && root.Val < q.Val { // 该树是个，查询树
			root = root.Right
		} else if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else {
			break
		}
	}
	return root
}

func lowestCommonAncestorV2(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 最近祖先节点，（最深）
	if root.Val < p.Val && root.Val < q.Val { // 该树是个，查询树
		return lowestCommonAncestorV2(root.Right, p, q)
	} else if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorV2(root.Left, p, q)
	}
	return root
}

func lowestCommonAncestorV3(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	max := int(math.Max(float64(p.Val), float64(q.Val)))
	min := int(math.Min(float64(p.Val), float64(q.Val)))
	for root != nil {
		if root.Val >= min && root.Val <= max {
			break
		} else if root.Val > max { // 该树是个，查询树
			root = root.Left
		} else if root.Val < min {
			root = root.Right
		}
	}

	return root
}
