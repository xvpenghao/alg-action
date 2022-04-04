package array_str

import (
	"fmt"
	"math"
)

// @Author XuPEngHao
// @DATE 2022/4/3 14:49

/*
剑指 Offer 57 - II. 和为s的连续正数序列
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。



示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]
*/

/*
数组是连续正数序列，
数组中至少含有2整数（比零大的数）
不同序列按照收个数字打下排列 [2,  3] [4  ,5] ,2<4

如果是奇数则可以缩小范围 [1,15/2+1]

如果是偶数则缩小范围[1-n/2]
6 [1,2,3]
10 [1,2,3,4]
等差数列公式 n*(a1+an)/2
*/

// 利用的滑动窗口，向右移动，控制左，右边界。
// 结束条件 i>j
func findContinuousSequence3(target int) [][]int {
	res := make([][]int, 0)
	left, right, sum := 1, 2, 3
	// target = 9 [[2,3,4],[4,5]]
	// [1,2,3,4,5]
	for left < right {
		if sum == target {
			res = append(res, getArray(left, right))
		}
		if sum >= target { // 向右移动，移动左边界（减少窗口内的元素和）减少
			sum -= left
			left++
		} else { // 向右移动，移动右边界（增大窗口内的元素和）增多元素
			right++
			sum += right
		}
	}
	return res
}

// todo 不理解
func findContinuousSequence2(target int) [][]int {
	res := make([][]int, 0)
	for i := int(math.Sqrt(float64(2 * target))); i >= 2; i-- {
		judge := target - i*(i-1)/2
		if judge%i == 0 {
			begin := judge / i
			temp := make([]int, 0)
			for j := 0; j < i; j++ {
				temp = append(temp, begin+j)
			}
			res = append(res, temp)
		}
	}
	return res
}

// 超时
func findContinuousSequence(target int) [][]int {
	if target < 2 {
		return nil
	}
	end := target / 2
	if target%2 != 0 {
		end += 1
	}

	// 输入：target = 9
	// [1,2,3,4,5]
	// 输出：[[2,3,4],[4,5]]

	// 输入：target = 15
	// [1,2,3,4,5,6,7,8]
	// 输出：[[1,2,3,4,5],[4,5,6],[7,8]]
	result := make([][]int, 0)
	n := 0
	var loop int
	for i := 1; i <= end-1; i++ {
		for j := end; j >= i+1; j-- {
			loop++
			n = (j - i + 1) * (i + j) / 2 // 发现小于内部循环直接break
			if n < target {
				break
			}
			if n == target {
				result = append(result, getArray(i, j))
				break
			}
		}
	}
	fmt.Println(loop)
	return result
}

func getArray(start, end int) []int {
	numList := make([]int, end-start+1)
	i := 0
	for start <= end {
		numList[i] = start
		i++
		start++
	}
	return numList
}
