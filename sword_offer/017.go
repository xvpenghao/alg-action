package sword_offer

import (
	"bytes"
	"math"
	"strconv"
)

/**
* @Author XuPEngHao
* @Date 2023/6/10 21:18
 */

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

// 3
// 1 2 3 4 5 6 7 8 9 10
// 11 12 13 14 15 16 17 18 19 20
// 21 22 23 24 25 26 27 28 29 30
// ... 200
// 201 202
// 字符串大数相加
func printNumbers2(n int) []string { // 考虑到打印大数的问题
	// base 是 1-10 每次加10来递增
	// 11 + 10
	result := make([]string, 0)
	res := "0"
	for {
		res = sumBigStr(res, "1")
		if len(res) > n {
			return result
		}
		result = append(result, res)
	}
}

// 模拟数字字符串相加
// 9 + 1 = 10
// 90 + 10 = 100
// 100
// 001
func sumBigStr(str1, str2 string) string {
	// 大数加+1
	overflow := 0
	str2Index := len(str2) - 1
	sumVal := make([]string, 0)
	for i := len(str1) - 1; i >= 0; i-- {
		str2Val := "0"
		if str2Index >= 0 {
			str2Val = string(str2[str2Index])
			str2Index--
		}
		str1Int, _ := strconv.Atoi(string(str1[i]))
		str2ValInt, _ := strconv.Atoi(str2Val)
		res := str1Int + str2ValInt + overflow
		if res == 10 { // 0 % 10 == 0
			overflow = 1
			res = 0
		} else {
			overflow = 0
		}
		sumVal = append(sumVal, strconv.Itoa(res))
	}

	if overflow > 0 {
		sumVal = append(sumVal, "1")
	}
	buf := bytes.Buffer{}
	for i := len(sumVal) - 1; i >= 0; i-- {
		buf.WriteString(sumVal[i])
	}
	return buf.String()
}

func printNumbers(n int) []int {
	size := int(math.Pow(10, float64(n))) - 1
	numsList := make([]int, size)
	for i := 1; i <= size; i++ {
		numsList[i-1] = i
	}
	return numsList
}
