package linked

// @Author XuPEngHao
// @DATE 2022/5/3 12:20

/*
剑指 Offer 24. 反转链表
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。



示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL


限制：

0 <= 节点个数 <= 5000
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseListV2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var h *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = h
		h = head
		head = tmp
	}
	return h
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 1-2-3-4-null
	// 2-next= 1
	// 1.next = 3
	// head = 2
	// head2 = t
	// 4-3-2-1-null
	var h *ListNode
	tail := head
	head2 := head.Next
	for head2 != nil {
		tmp := head2.Next

		tail.Next = head2.Next
		head2.Next = h
		h = head2

		tail = tail.Next // 下个尾节点
		head2 = tmp
	}
	return h
}
