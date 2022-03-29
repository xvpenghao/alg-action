package array_str

import "sort"

// @Author XuPEngHao
// @DATE 2022/3/27 12:21

/*
剑指 Offer 39. 数组中出现次数超过一半的数字
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。



你可以假设数组是非空的，并且给定的数组总是存在多数元素。



示例 1:

输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2


限制：

1 <= 数组长度 <= 50000
*/

// 摩尔投票法，相互抵消
func majorityElementV4(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 直接不用计算
	if len(nums) == 1 {
		return nums[0]
	}
	x, votes := 0, 0
	// {8, 8, 7, 7, 7}
	for _, val := range nums {
		if votes == 0 {
			x = val
		}
		if val == x {
			votes++
		} else {
			votes--
		}
	}

	// 如果说，众数不一定存在，则遍历 nums数组计算 nums[i]==x出现的次数 > len(nums)>1 即可
	return x
}

// hashmap
func majorityElementV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 直接不用计算
	if len(nums) == 1 {
		return nums[0]
	}
	halfNumLen := len(nums) >> 1
	numMap := make(map[int]int, 50000>>1)
	cnt, ok := 0, false
	for _, val := range nums {
		cnt, ok = numMap[val]
		if !ok {
			numMap[val] = 1
			continue
		}
		if cnt+1 > halfNumLen {
			return val
		}
		numMap[val] = cnt + 1
	}
	return 0
}

// 数组中有一个数字出现的次数超过数组长度的一半
// 排序完成后，直接去中间的就行了
func majorityElementV3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 直接不用计算
	if len(nums) == 1 {
		return nums[0]
	}
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[len(nums)>>1]
}

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 直接不用计算
	if len(nums) == 1 {
		return nums[0]
	}
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	halfNumLen := len(nums) >> 1
	ctn := 1
	// [1, 2, 3, 2, 2, 2, 5, 4, 2]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			ctn++
		} else {
			ctn = 1
		}
		if ctn > halfNumLen {
			return nums[i]
		}
	}
	return 0
}
