package stack_queue

import "testing"

// @Author XuPEngHao
// @DATE 2022/3/10 19:54

func Test_CQueue(t *testing.T) {
	cq := CQueue2{}
	cq.AppendTail(1)
	cq.AppendTail(2)
	cq.AppendTail(3)

	t.Log(cq.DeleteHead())
	t.Log(cq.DeleteHead())
	t.Log(cq.DeleteHead())
	t.Log(cq.DeleteHead())

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
	t.Log(reversePrintV3(head))
}
