package array_str

// @Author XuPEngHao
// @DATE 2022/4/1 08:57

/*
剑指 Offer 53 - II. 0～n-1中缺失的数字
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。



示例 1:

输入: [0,1,3]
输出: 2
示例 2:

输入: [0,1,2,3,4,5,6,7,9]
输出: 8
*/

// 数组的元素是，[0,n]，长度[0,n)
// 数组元素的一个，递增的连续性的，
// 核心的问题知识点是 nums[i]=i
func missingNumberV2(nums []int) int {
	if nums[0] == 1 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	// [0,1,2]
	return len(nums)
}

// 递增排序数组
// 唯一的 && 连续的
// 左边一定是连续的，则继续mid+1
func missingNumber(nums []int) int {
	// [0,1,3,4,5] // 2
	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		mid = (left + right) >> 1
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
