package array_str

// @Author XuPEngHao
// @DATE 2022/3/16 17:39

/*
剑指 Offer 11. 旋转数组的最小数字
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。

给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。请返回旋转数组的最小元素。
例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为1。

示例 1：

输入：[3,4,5,1,2] [6,1,2,3,4,5]
输出：1
示例 2：

输入：[2,2,2,0,1]
输出：0
*/

func minArrayV2(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	startIndex, endIndex := 0, len(numbers)-1
	m := 0
	for startIndex < endIndex {
		m = (startIndex + endIndex) / 2
		if numbers[m] > numbers[endIndex] {
			startIndex = m + 1
		} else if numbers[m] < numbers[endIndex] {
			endIndex = m
		} else {
			endIndex--
		}
	}
	return numbers[startIndex]
}

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}

	endIndex := len(numbers) - 1
	minNum := numbers[endIndex] // 最后一个
	endIndex--
	for endIndex >= 0 {
		if numbers[endIndex] > minNum {
			return minNum
		}
		minNum = numbers[endIndex]
		endIndex--
	}
	return minNum
}
