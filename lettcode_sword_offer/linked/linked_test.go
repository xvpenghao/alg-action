package linked

import "testing"

// @Author XuPEngHao
// @DATE 2022/5/3 10:54

func Test_deleteNode(t *testing.T) {
	node := deleteNodeV2(&ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}, 5)
	t.Log(node)
}
