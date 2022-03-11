package stack_queue

// @Author XuPEngHao
// @DATE 2022/3/10 19:42

/*
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
示例 1：
输入：head = [1,3,2]
输出：[2,3,1]

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrintV3(head *ListNode) []int {
	if head == nil {
		return nil
	}
	head2 := head.Next
	head.Next = nil // 尾部就要nil
	loop := 1
	for head2 != nil {
		tmp := head2.Next
		head2.Next = head
		head = head2
		head2 = tmp
		loop++
	}
	nums, i := make([]int, loop), 0
	for head != nil {
		nums[i] = head.Val
		head = head.Next
		i++
	}
	return nums
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var numList []int
	for head != nil {
		numList = append(numList, head.Val)
		head = head.Next
	}
	numLen := len(numList)
	result := make([]int, numLen)

	for index := range numList {
		result[index] = numList[numLen-index-1]
	}
	return result
}
