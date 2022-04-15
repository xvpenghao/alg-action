package stack_queue

import "sort"

// @Author XuPEngHao
// @DATE 2022/4/12 08:21

/*
剑指 Offer 30. 包含min函数的栈
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。


示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.
*/

/*
当时的难点是，我每次的入栈，可能都会是最小的元素。如何不需要排序，就能记住最下的元素呢？
那就记录每次入栈的最小值，这样 minStack，里面都是 从大到小的的栈元素
【-2,1,-3,4,-9】
【-2,-3,-9】*/
type MinStack2 struct {
	stack     []int
	helpStack []int // 腐竹栈
}

func ConstructorMinStack2() MinStack2 {
	return MinStack2{}
}

func (this *MinStack2) Push(x int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, x)
		this.helpStack = append(this.helpStack, x)
		return
	}
	this.stack = append(this.stack, x)
	// 新加入的元素 和栈顶，元素比较，比辅助栈顶元素小，则加入栈
	if x <= this.helpStack[len(this.helpStack)-1] {
		this.helpStack = append(this.helpStack, x)
	}
}

func (this *MinStack2) Pop() {
	if len(this.stack) == 0 {
		return
	}
	// 判断弹出的 栈顶的元素 是否和 辅助栈顶相等，相等则同样弹出
	popEle := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if popEle != this.helpStack[len(this.helpStack)-1] {
		return
	}
	this.helpStack = this.helpStack[:len(this.helpStack)-1]
}

func (this *MinStack2) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack2) Min() int {
	if len(this.helpStack) == 0 {
		return 0
	}
	return this.helpStack[len(this.helpStack)-1]
}

type MinStack struct {
	stack       []int
	top         int
	sortEleList []int
	m           map[int]int // 假设值不会重复
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	return MinStack{
		m: make(map[int]int),
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.top = x

	this.sortEleList = append(this.sortEleList, x)
	sort.SliceStable(this.sortEleList, func(i, j int) bool {
		return this.sortEleList[i] < this.sortEleList[j]
	})
	for idx, val := range this.sortEleList {
		this.m[val] = idx
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	popEle := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if len(this.stack) > 0 {
		this.top = this.stack[len(this.stack)-1]
	}
	// 当pop完毕后，长度=0，则top 和 min无值
	if len(this.stack) == 0 {
		this.top = 0
		this.m = make(map[int]int)
		this.sortEleList = nil
		return
	}

	// 删除 m和sortEle
	delIdx := this.m[popEle]
	delete(this.m, popEle)
	if delIdx == 0 {
		this.sortEleList = this.sortEleList[1:]
	} else if delIdx == len(this.sortEleList)-1 {
		this.sortEleList = this.sortEleList[:len(this.sortEleList)-1]
	} else {
		this.sortEleList = append(append([]int{}, this.sortEleList[:delIdx]...), this.sortEleList[delIdx+1:]...)
	}
	for idx, val := range this.sortEleList {
		this.m[val] = idx
	}
}

func (this *MinStack) Top() int {
	return this.top
}

func (this *MinStack) Min() int {
	if len(this.sortEleList) == 0 {
		return 0
	}
	return this.sortEleList[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
