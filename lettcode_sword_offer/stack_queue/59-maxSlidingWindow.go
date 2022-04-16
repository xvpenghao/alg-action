package stack_queue

// @Author XuPEngHao
// @DATE 2022/4/15 09:46

/*
剑指 Offer 59 - I. 滑动窗口的最大值
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----  //
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


提示：

你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
*/

/*

[1,2,3,4] 1 4次   1,2,3,4
[1,2,3,4] 2 3次   1,2 | 2,3 | 3,4
[1,2,3,4] 3 2次  1,2,3 | 2,3,4
[1,3,-1,-3,5,3,6,7] 3 6次
*/

// 是否能提前算出来，能够分几次呢？
// 滑动窗口的中数字最大计算，是否可以优化

// 始终保证队列，单调递减，并且队列中始终是，滑动窗口的的值，保持新鲜
// 你不能用max一个变量来及，因为他可能别边缘值，被提出，那你就用一个queue来记，最大，二大，次da，单调递减
func maxSlidingWindowV2(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}

	queue := make([]int, 0)  // 保持单调递减
	for i := 0; i < k; i++ { //
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}

	result := make([]int, len(nums)-k+1)
	result[0] = queue[0]
	for i := k; i < len(nums); i++ {
		if len(queue) > 0 && queue[0] == nums[i-k] { // 被移出的那一个如果等于 queue=0，则弹窗。1,3,-1,-3,5,3,6,7
			queue = queue[1:]
		}
		// 保持递减
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		result[i-k+1] = queue[0]
	}
	return result
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}

	result := make([]int, len(nums)-k+1)
	index := 0
	// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
	for i := 0; i < len(nums)-k+1; i++ {
		result[index] = getMaxNum(nums[i : i+k])
		index++
	}
	return result
}

func getMaxNum(numList []int) int {
	max := numList[0]
	for i := 1; i < len(numList); i++ {
		if numList[i] > max {
			max = numList[i]
		}
	}
	return max
}
