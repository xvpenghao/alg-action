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

/*
0、头部节点的下一个为nil
0、头部节点遍历
0、头插发
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var h *ListNode
	for head != nil {
		tmp := head.Next // 保存一个节点
		head.Next = h    // 头插发
		h = head         // 重新赋值h
		head = tmp       // 头节点下一个赋值给 head，接着循环遍历
	}
	return h
}
