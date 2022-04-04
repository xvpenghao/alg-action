package array_str

// @Author XuPEngHao
// @DATE 2022/4/4 15:19

/*
剑指 Offer 57. 和为s的两个数字
输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。



示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[2,7] 或者 [7,2]
示例 2：

输入：nums = [10,26,30,31,47,60], target = 40
输出：[10,30] 或者 [30,10]

[1,2,3,4,5,6]  11,
*/

// 对撞双指针
// 利用数组的升序的特性
func twoSum3(nums []int, target int) []int {
	nums = nums[:getInsertLoc(nums, target)]
	left, right := 0, len(nums)-1
	sum := 0
	for left < right {
		sum = nums[left] + nums[right]
		if sum == target {
			return []int{nums[left], nums[right]}
		} else if sum > target { // 因为该数组是升序的，所以缩小右边的范围
			right--
		} else {
			left++
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	nums = nums[:getInsertLoc(nums, target)]
	numsMap := make(map[int]struct{}, len(nums))
	for _, val := range nums {
		if _, ok := numsMap[target-val]; ok {
			return []int{val, target - val}
		}
		numsMap[val] = struct{}{}
	}
	return nil
}

// 在有序数组中，找出个数使其他们的和为 target
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	// 目的是缩小搜索的范围
	nums = nums[:getInsertLoc(nums, target)]
	if len(nums) == 2 && nums[0]+nums[1] == target {
		return nums
	}
	for index, val := range nums {
		if binarySearch3(nums[index:], target-val) > 0 {
			return []int{val, target - val}
		}
	}
	return nil
}

func binarySearch3(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		mid = (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0
}

// [2,7,11,15], target = 9
func getInsertLoc(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		mid = (left + right) >> 1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
