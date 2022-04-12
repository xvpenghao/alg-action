package stack_queue

// @Author XuPEngHao
// @DATE 2022/4/9 20:33

/*
剑指 Offer 09. 用两个栈实现队列
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )



示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用
*/

/*
栈：先进后出
栈：先进后出
队列：先进先出
*/

type CQueue2 struct {
	tailStack []int
	headStack []int
}

func Constructor2() CQueue2 {
	return CQueue2{}
}

func (this *CQueue2) AppendTail(value int) {
	this.tailStack = append(this.tailStack, value)
}

func (this *CQueue2) DeleteHead() int {
	if len(this.headStack) > 0 { // 如果头栈有元素则直接弹出
		popVal := this.headStack[len(this.headStack)-1]
		this.headStack = this.headStack[:len(this.headStack)-1]
		return popVal
	}

	if len(this.tailStack) == 0 { // 没有可用的元素
		return -1
	}

	// 将 tailStack中数据pop 到HeadStack
	this.headStack = make([]int, len(this.tailStack))
	for i := len(this.tailStack) - 1; i >= 0; i-- {
		this.headStack[len(this.tailStack)-1-i] = this.tailStack[i]
	}
	this.tailStack = nil // pop完毕后 TailStack就为nil了
	popEle := this.headStack[len(this.headStack)-1]
	this.headStack = this.headStack[:len(this.headStack)-1]
	return popEle
}

type CQueue struct {
	tailStack []int
	headStack []int
	size      int // 栈中的元素
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.tailStack = append(this.tailStack, value)
	this.size++
}

func (this *CQueue) DeleteHead() int {
	if this.size < 1 {
		return -1
	}
	// 将 tailStack中数据pop 到HeadStack
	this.headStack = make([]int, this.size)
	for i := this.size - 1; i >= 0; i-- {
		this.headStack[this.size-1-i] = this.tailStack[i]
	}
	this.tailStack = nil // pop完毕后 TailStack就为nil了

	// 弹窗元素
	popEle := this.headStack[this.size-1]
	this.headStack = this.headStack[:this.size-1]
	this.size--

	// 在放到inStack中
	this.tailStack = make([]int, this.size)
	for i := this.size - 1; i >= 0; i-- {
		this.tailStack[this.size-1-i] = this.headStack[i]
	}
	this.headStack = nil
	return popEle
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
