package sword_offer

import "sort"

/**
* @Author XuPEngHao
* @Date 2023/6/5 10:55
 */

/*
找出数组中重复的数字。

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3

限制： 2 <= n <= 100000
*/

//  在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内
// 完成交换-是指 下标和元素之相等
// 0 1 2 3
// 0 3 0 1
func findRepeatNumberV3(nums []int) int { // 当下标和下标相等则i++
	// 映射
	i := 0
	for i < len(nums) {
		if i == nums[i] { // 相等则无需交互
			i++
			continue
		}
		if nums[nums[i]] == nums[i] { // 数字下标值 == 此时下标值
			return nums[i]
		}
		// 下标值 和 值下标值交互 核心的交换
		numIndex := nums[i]
		nums[numIndex], nums[i] = nums[i], nums[numIndex]
	}
	return -1
}

func findRepeatNumberV2(nums []int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

func findRepeatNumber(nums []int) int {
	numMap := make(map[int]int, len(nums)) // 使用空间换时间的
	for _, n := range nums {
		val, ok := numMap[n]
		if ok {
			return n // 返回的数字
		}
		numMap[n] = val + 1
	}
	return -1
}
