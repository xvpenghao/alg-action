package array_str

// @Author XuPEngHao
// @DATE 2022/3/18 09:21

/*
剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。



示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。


提示：

0 <= nums.length <= 50000
0 <= nums[i] <= 10000
*/

// 利用快排的特点，左走，右走的特点
// 要学会善于发现 命题中 描述的一些特点
func exchangeV3(numList []int) []int {
	if len(numList) == 0 {
		return nil
	}
	startIndex, endIndex := 0, len(numList)-1
	for startIndex < endIndex {
		for startIndex < endIndex && numList[startIndex]&1 == 1 { // 如果左边是奇数则 startIndex++
			startIndex++
		}
		for startIndex < endIndex && numList[endIndex]&1 == 0 { // 如果右边是偶数，则 endIndex--
			endIndex--
		}
		// startIndex 不等于 endIndex则交换位置
		if startIndex != endIndex {
			numList[startIndex], numList[endIndex] = numList[endIndex], numList[startIndex]
		}
	}
	return numList
}

func exchangeV2(numList []int) []int {
	if len(numList) == 0 {
		return nil
	}
	startIndex, endIndex := 0, len(numList)-1
	startJ, endO := false, false
	for startIndex < endIndex { // [1,2,3,4]
		// 如果 startIndex 是偶数
		// 如果 endIndex 是奇数 则交换
		startJ = numList[startIndex]%2 != 0
		endO = numList[endIndex]%2 == 0
		if !startJ && !endO { // 如果startIndex偶数，endIndex是奇数则交换
			numList[startIndex], numList[endIndex] = numList[endIndex], numList[startIndex]
			startIndex++
			endIndex--
			continue
		}
		if startJ { // startIndex 是奇数
			startIndex++
		}
		if endO { // startIndex 是偶数
			endIndex--
		}
	}
	return numList
}

// 浪费了空间，节省了时间
func exchange(numList []int) []int {
	if len(numList) == 0 {
		return nil
	}

	newNumList := make([]int, len(numList))
	startIndex, endIndex := 0, len(numList)-1
	for _, val := range numList {
		if val%2 == 0 { // 如果是偶数则放大后面
			newNumList[endIndex] = val
			endIndex--
			continue
		}
		newNumList[startIndex] = val // 奇数则放到前面
		startIndex++
	}
	return newNumList
}
