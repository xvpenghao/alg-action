package linked

// @Author XuPEngHao
// @DATE 2022/5/4 15:14

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	A, B := headA, headB
	for A != B {
		if A != nil {
			A = A.Next
		} else {
			A = headB
		}

		if B != nil {
			B = B.Next
		} else {
			B = headA
		}
	}
	return A
}
