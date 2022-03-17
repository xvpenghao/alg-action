package array_str

import "math"

// @Author XuPEngHao
// @DATE 2022/3/17 09:17
/*
剑指 Offer 17. 打印从1到最大的n位数
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:

输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]


说明：

用返回一个整数列表来代替打印
n 为正整数
*/

/*
1- 10^1-1,
2- 10^2-1
3- 10^3-1
*/
func printNumbers(n int) []int {
	if n < 1 {
		return nil
	}
	end := int(math.Pow(10, float64(n))) - 1
	numList := make([]int, end)
	index, start := 0, 1
	for index < end {
		numList[index] = start
		start++
		index++
	}
	return numList
}
