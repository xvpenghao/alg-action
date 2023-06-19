package sword_offer

/**
* @Author XuPEngHao
* @Date 2023/6/18 18:41
 */

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
*/

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}

	// 重新模拟一遍的 push的弹窗不就行了吗
	stackObj := make([]int, 0)
	poppedStart := 0
	for _, pushVal := range pushed {
		stackObj = append(stackObj, pushVal)
		// 栈中元素是否和 popped 的元素相等
		for len(stackObj) > 0 {
			if stackObj[len(stackObj)-1] == popped[poppedStart] {
				stackObj = stackObj[:len(stackObj)-1]
				poppedStart++
			} else {
				break
			}
		}
	}
	// 使用栈来做最后的判断结果
	return len(stackObj) == 0
}
