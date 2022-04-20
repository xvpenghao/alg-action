package hash_table

import (
	"sort"
)

// @Author XuPEngHao
// @DATE 2022/4/19 09:02

/*
剑指 Offer 03. 数组中重复的数字
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入： 一个萝卜一个坑位。
[2, 3, 1, 0, 2, 5, 3]

[0,1,2,3,3,4,5]
输出：2 或 3


限制：

2 <= n <= 100000
*/

// 认真读题，发现题中的关键信息[0,n-1]
func findRepeatNumberV3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 根据下标交换坑位 将元素交换到正确的下标位置，然后在忽略他
	index := 0
	for index < len(nums) {
		if index == nums[index] {
			index++
			continue
		}
		if nums[index] == nums[nums[index]] {
			return nums[index]
		}
		nums[nums[index]], nums[index] = nums[index], nums[nums[index]]
	}
	return -1
}

func findRepeatNumberV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	index := 0
	for index+1 < len(nums) {
		if nums[index] == nums[index+1] {
			return nums[index]
		}
		index++
	}
	return 0
}

func findRepeatNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]struct{}, len(nums))
	for _, val := range nums {
		if _, ok := m[val]; ok {
			return val
		}
		m[val] = struct{}{}
	}
	return 0
}
