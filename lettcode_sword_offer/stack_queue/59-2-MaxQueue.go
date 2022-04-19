package stack_queue

// @Author XuPEngHao
// @DATE 2022/4/16 20:33

/*
剑指 Offer 59 - II. 队列的最大值
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]


限制：

1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5
*/

// [3,7,5,2,10]
// 单吊递减
// 3, [3] , 7 [7], 5 [7,5]  2  [7,5,2] 10 [10]

// [10,9,8,2,1]

// 10,[10], 9

type MaxQueue struct {
	queue   []int // 队列
	dequeue []int // 双端队列。 可以在两端进行插入和删除
}

func ConstructorMaxQueue() MaxQueue {

	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.dequeue) == 0 {
		return -1
	}
	return this.dequeue[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	if len(this.dequeue) == 0 {
		this.dequeue = append(this.dequeue, value)
		return
	}
	// 保持双端队列单调递减
	for len(this.dequeue) > 0 && this.dequeue[len(this.dequeue)-1] < value {
		this.dequeue = this.dequeue[:len(this.dequeue)-1]
	}
	this.dequeue = append(this.dequeue, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	popVal := this.queue[0]
	this.queue = this.queue[1:]
	if len(this.dequeue) > 0 && this.dequeue[0] == popVal {
		this.dequeue = this.dequeue[1:]
	}
	return popVal
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
