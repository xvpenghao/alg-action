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

/*
0、首先分析该 旋转数组的特性，
0、以旋转点为分界线，则两边是有序的，
0、从有序数组里面查询数值，常用的查找算法就是 【二分查找算法】
0、核心就是，找到有序的数组的分段，使用二分查找算法找到最小的值
*/
func minArrayV3(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	// [4.5,1,2,3]
	start, end := 0, len(numbers)-1
	mid := 0
	for start < end {
		mid = (start + end) >> 1
		// 判断中间的值，来缩小判定的范围
		if numbers[mid] > numbers[end] { // 说明 m在一定在左边，旋转点在[m+1,end]，m可定不是旋转点
			start = mid + 1
		} else if numbers[mid] < numbers[end] { // 说明 m一定在 右边 或者本身就是旋转点，旋转点在 [i,m]
			end = mid
		} else { // numbers[mid] == numbers[end]  // m在右边[3,1,3,3,3]  m在左边 [3,3,1,3] 不好，判断则所以end的范围
			end--
		}
	}
	return numbers[start]
}

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
