package array_str

// @Author XuPEngHao
// @DATE 2022/3/30 08:36

/*
剑指 Offer 53 - I. 在排序数组中查找数字 I
统计一个数字在排序数组中出现的次数。



示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0


提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109
*/

// 找出左边界（target-1） 和右边界，然后用左边界- 右边界 即可
func searchV3(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 && nums[0] == target {
		return target
	}
	// 插入target插入的合适位置，每次找到都是相同的最大的位置
	return getTargetInsertIndex(nums, target) - getTargetInsertIndex(nums, target-1)
}

// 找出元素合适的插入位置（二分查找发，可以演变为 二分插入发，每次返回一次合适的插入位置）
func getTargetInsertIndex(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	// [1,2,8,8,9]
	for left <= right {
		mid = (right + left) >> 1
		if nums[mid] <= target { // 最后的返回 left ，很大关系就和这个有关
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func searchV2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 && nums[0] == target {
		return target
	}
	targetCnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			targetCnt++
		}
	}
	return targetCnt
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 && nums[0] == target {
		return target
	}

	// 二分查找，返回查询到下标
	targetIndex := binarySearch2(nums, target)
	if targetIndex < 0 {
		return 0
	}
	targetCnt := 1
	// 左走
	for i := targetIndex - 1; i >= 0; i-- {
		if nums[i] != target {
			break
		}
		targetCnt++
	}
	// 右走
	for i := targetIndex + 1; i < len(nums); i++ {
		if nums[i] != target {
			break
		}
		targetCnt++
	}
	return targetCnt
}

// nums是一个升序的数组
func binarySearch2(nums []int, target int) int {
	// [2,2],target=3, 你就会发现  len(nums)-1
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

	return - 1
}
