package stack_queue

import "testing"

// @Author XuPEngHao
// @DATE 2022/3/10 19:54

func Test_ReversePrintV2(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	t.Log(reversePrintV3(head))
}
func Test_ReversePrint(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	t.Log(reversePrint(head))
}
