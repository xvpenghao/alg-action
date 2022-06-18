package recursion

// @Author XuPEngHao
// @DATE 2022/6/1 09:12

/*
剑指 Offer 26. 树的子结构
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

     3
    / \
   4   5
  / \
 1   2
给定的树 B：

   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：

输入：A = [1,2,3], B = [3,1]
输出：false
示例 2：

输入：A = [3,4,5,1,2], B = [4,1]
输出：true
限制：
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubStructureV2(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (recur(A, B) || isSubStructureV2(A.Left, B) || isSubStructureV2(A.Right, B))
}

// 比较两个树是否相等
// 递归的使用思想，确定【终止条件】【返回条件循环添加】
func recur(A *TreeNode, B *TreeNode) bool {
	if B == nil { // B已经用尽
		return true
	}
	if A == nil || (A.Val != B.Val) { // A用尽 或者 两者的值不相等则直接返回false
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}

// 节点的左右页需要对上。
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	// 找到符号的子树吗，知道没有更合适的节点
	A = findNodeByTarget(A, B.Val)
	if A == nil {
		return false
	}
	if dfsEqual(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func dfsEqual(A *TreeNode, B *TreeNode) bool {
	// 不相等立马退出
	if A != nil && B != nil && A.Val != B.Val {
		return false
	}
	if A == nil && B != nil {
		return false
	}
	if A != nil && B != nil {
		return dfsEqual(A.Left, B.Left) && dfsEqual(A.Right, B.Right)
	}
	return true
}

// 找打要比较的节点位置
func findNodeByTarget(A *TreeNode, targetNum int) *TreeNode {
	if A == nil {
		return nil
	}
	if A.Val == targetNum {
		return A
	}
	if node := findNodeByTarget(A.Left, targetNum); node != nil {
		return node
	}
	if node := findNodeByTarget(A.Right, targetNum); node != nil {
		return node
	}
	return nil
}
