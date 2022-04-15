package stack_queue

// @Author XuPEngHao
// @DATE 2022/4/12 10:18

/*
剑指 Offer 31. 栈的压入、弹出序列
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。



示例 1：

输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
示例 2：

输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。


提示：

0 <= pushed.length == popped.length <= 1000
0 <= pushed[i], popped[i] < 1000
pushed 是 popped 的排列。
*/
/*
push [1,2,3,4,5]
pop [4,5,3,2,1]
pop [1,2,3,4,5]
pop [5,4,3,2,1]
pop [4,5,3,2,1]
pop [4,3,5,2,1]
pop [4,3,5,1,2] false
*/
// 利用的栈的特性的进行操作
// 可以分析一下，他是怎么把 push 和pop 划分为2个数组，模式执行流程
func validateStackSequences2(pushed []int, popped []int) bool {
	var stack []int
	pushedIdx, poppedIdx := 0, 0
	for pushedIdx < len(pushed) {
		stack = append(stack, pushed[pushedIdx])
		// 判断栈中的元素是否 popped的数据相等
		for len(stack) > 0 && stack[len(stack)-1] == popped[poppedIdx] {
			stack = stack[:len(stack)-1]
			poppedIdx++
		}
		pushedIdx++
	}
	return len(stack) == 0 // 当栈中的元素全部弹出完毕则表示相等
}

// 我以上来吧， pushed 和popped 两者强制的拉上关系了，并没有真正的去分析这个，答案（两个数组的怎么产生的）的来源
func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}

	if len(pushed) == 0 || len(popped) == 0 {
		return false
	}
	// pushed，摸你入栈 和 popped的元素对比，
	// false的情况是， 当前对比的元素的不相等，&& （pushed没有可用push元素了， 或者 push的元素不相等）则返回false
	poppedIndex := 0
	popEle := popped[poppedIndex]

	// 首先往栈中push一个元素。则 pushed的个数减少
	stack := make([]int, 0, len(pushed))
	stack = append(stack, pushed[0])
	pushed = pushed[1:]

	for len(stack) > 0 || len(pushed) > 0 {
		popEle = popped[poppedIndex]
		if len(stack) == 0 || popEle != stack[len(stack)-1] { // 判断是否相等于 stack的栈顶元素
			if len(pushed) > 0 { // push
				stack = append(stack, pushed[0]) // 往栈中添加元素
				pushed = pushed[1:]              // push元素减少一个
			} else {
				return false
			}
			continue
		}
		stack = stack[:len(stack)-1] // 相等则弹出元素
		poppedIndex++
	}

	return true
}
