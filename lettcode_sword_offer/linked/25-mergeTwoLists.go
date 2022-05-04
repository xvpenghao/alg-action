package linked

// @Author XuPEngHao
// @DATE 2022/5/4 12:28

/*
剑指 Offer 25. 合并两个排序的链表
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：

0 <= 链表长度 <= 1000
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoListsV2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	cur := new(ListNode)
	dump := cur
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	cur.Next = l1
	if l2 != nil {
		cur.Next = l2
	}
	return dump.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 1,4,10,11
	// 2,3,8

	h, h2, pre := l1, l2, l1
	if l2.Val < l1.Val {
		h, pre, h2 = l2, l2, l1
	}
	res := h
	for h != nil {
		for h2 != nil && h2.Val < h.Val { // 连续性的插入
			tmp := h2.Next // 保存下一个节点
			h2.Next = h
			pre.Next = h2
			pre = h2 // 设置pre
			h2 = tmp
		}
		pre = h
		h = h.Next
	}

	if h2 != nil {
		pre.Next = h2
	}

	return res
}
