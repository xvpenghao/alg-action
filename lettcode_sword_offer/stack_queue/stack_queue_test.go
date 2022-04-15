package stack_queue

import "testing"

// @Author XuPEngHao
// @DATE 2022/3/10 19:54

func Test_validateStackSequences(t *testing.T) {
	// t.Log(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	t.Log(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
	// t.Log(validateStackSequences([]int{4, 0, 1, 2, 3}, []int{4, 2, 3, 0, 1}))
	// t.Log(validateStackSequences([]int{1, 0}, []int{1, 0}))
}

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
