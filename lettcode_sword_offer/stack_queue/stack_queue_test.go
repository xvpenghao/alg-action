package stack_queue

import "testing"

// @Author XuPEngHao
// @DATE 2022/3/10 19:54

func Test_MinStack(t *testing.T) {
	ms := ConstructorMinStack2()
	ms.Push(2)
	ms.Push(0)
	ms.Push(3)
	ms.Push(0)

	// [2,0]
	t.Log("Min", ms.Min())
	ms.Pop()

	t.Log("Min", ms.Min())
	ms.Pop()

	t.Log("Min", ms.Min())
	ms.Pop()

	t.Log("Min", ms.Min())
}

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
