package linked

// @Author XuPEngHao
// @DATE 2022/5/3 10:54

/*
剑指 Offer 18. 删除链表的节点
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点。

注意：此题对比原题有改动

示例 1:

输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
示例 2:

输入: head = [4,5,1,9], val = 1
输出: [4,5,9]
解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.


说明：

题目保证链表中节点的值互不相同
若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNodeV2(head *ListNode, val int) *ListNode {
	if head == nil { // 链表为空 或者  要删除链表的头节点
		return nil
	}
	// 如果删除节点是投节点
	if head.Val == val {
		return head.Next
	}

	// 删除节点值，返回链表
	pre, cur := head, head.Next
	for cur != nil && cur.Val != val { // 找到要删除的节点
		pre = cur
		cur = cur.Next
	}
	if cur != nil {
		pre.Next = cur.Next // 节点删除
	}
	return head
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil { // 链表为空 或者  要删除链表的头节点
		return nil
	}
	// 如果删除节点是投节点
	if head.Val == val {
		head = head.Next
		return head
	}

	// 删除节点值，返回链表
	front := head
	nextHead := head.Next
	for nextHead != nil {
		if nextHead.Val == val {
			front.Next = nextHead.Next
			break
		}
		front = nextHead
		nextHead = nextHead.Next
	}
	return head
}
