package sword_offer

import "fmt"

/**
* @Author XuPEngHao
* @Date 2023/6/14 21:53
 */

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

/*
1，2，3，,4
1 4

1 5,2,4
*/

// 1 2 3  1 3 2
func exchangeV2(nums []int) []int {
	// 正序查找如果是奇数则不动，如果是偶数的倒着找到一个奇数来交换
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[i]%2 > 0 { // 奇数
			i++
		}
		if i < j && nums[j]%2 == 0 { // 偶数
			j--
		}
		if i != j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}

// 奇数比偶数大  奇数在前面
func exchange(nums []int) []int {
	quickSortV2(nums, 0, len(nums)-1)
	return nums
}

func quickSortV2(numList []int, left, right int) {
	if left < right {
		i, j := left, right
		tmp := numList[i]
		for i < j {
			// 左边大 右边小 左奇右偶
			for i < j && numList[j]%2 <= tmp%2 {
				j--
			}
			if i < j {
				numList[i] = numList[j]
				i++
			}
			// 从左边
			for i < j && numList[i]%2 > tmp%2 {
				i++
			}
			if i < j {
				numList[j] = numList[i]
				j--
			}
		}
		numList[j] = tmp
		quickSortV2(numList, left, i-1) // 递归调用
		quickSortV2(numList, i+1, right)
	}
}

func quickSort(numList []int, left, right int) {
	if left < right {
		i := qSort(numList, left, right)
		qSort(numList, left, i-1) // 递归调用
		qSort(numList, i+1, right)
	}
}

// 退出是返回i
func qSort(numList []int, left, right int) int {
	tmp := numList[left]

	for left < right {
		// 左边大 右边小 左奇右偶
		for left < right && numList[right]%2 < tmp%2 {
			fmt.Println(numList[right])
			right--
		}
		if left < right {
			fmt.Println(numList[right])
			numList[left] = numList[right]
			left++
		}
		// 从左边
		for left < right && numList[left]%2 >= tmp%2 {
			left++
		}
		if left < right {
			numList[right] = numList[left]
			right--
		}
	}
	numList[left] = tmp
	return left
}
