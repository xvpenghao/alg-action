package hot_100

/**
* @Author XuPEngHao
* @Date 2023/6/9 10:57
 */

/*
1. 两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。



示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
*/

// 不要想的太复杂了
func twoSum2(nums []int, target int) []int {
	numIndexMap := make(map[int]int, 0)
	for i, num := range nums {
		if index, ok := numIndexMap[target-num]; ok {
			return []int{index, i}
		}
		numIndexMap[num] = i
	}
	return nil
}

// 不需要将都一上来都填充到map中
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int][]int, 0)
	for index, num := range nums {
		numsMap[num] = append(numsMap[num], index)
	}

	for i, n := range nums {
		n2 := target - n
		if n == n2 && len(numsMap[n]) >= 2 { // 重复的2个数字
			return numsMap[n][:2]
		}

		if indexList, ok := numsMap[n2]; ok && indexList[0] != i {
			return []int{i, indexList[0]}
		}
	}

	return nil
}
