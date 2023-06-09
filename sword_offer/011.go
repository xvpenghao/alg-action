package sword_offer

/**
* @Author XuPEngHao
* @Date 2023/6/7 21:47
 */
/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。

给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。请返回旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为 1。  

注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。


示例 1：

输入：numbers = [3,4,5,1,2]
输出：1
示例 2：

输入：numbers = [2,2,2,0,1]
输出：0
*/

// 参考大佬实例在学习一下

// 利用二分查找的思想来思考问题
// 1,2,3,4,5 之前是个升序数组
// 3,4,5,1,2 旋转后的,

// 最慢的
func minArray3(numbers []int) int {
	for i := range numbers {
		if i < len(numbers)-1 && numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}
	return numbers[0]
}

func minArray2(numbers []int) int {
	i, j := 0, len(numbers)-1
	for i < j {
		mid := (i + j) >> 1
		if numbers[mid] > numbers[j] { // Mid 一定在 左排序数组中, 旋转点我们让他在右边找 [mid+1,j]
			i = mid + 1
		} else if numbers[mid] < numbers[j] { // Mid 一定在有排序数组中，旋转点我们让他在左边找 [i,m]
			j = mid
		} else {
			x := i
			for k := i + 1; k < j; j++ {
				if numbers[k] < numbers[x] {
					x = k
				}
			}
			return numbers[x]
		}
	}
	return numbers[i]
}

// 1,2,3,4,5 之前是个升序数组
// 3,4,5,1,2 旋转后的,
func minArray(numbers []int) int {
	// 判断是否是旋转数组 ，一旦产生旋转则， nums[0] > nums[end]
	// 否则则不是旋转数组
	if numbers[0] < numbers[len(numbers)-1] {
		return numbers[0]
	}

	min := numbers[len(numbers)-1]
	for i := len(numbers) - 1; i > 0; i-- {
		if numbers[i-1] <= numbers[i] { // 倒序找出分割线
			continue
		}
		return numbers[i]
	}
	return min
}
