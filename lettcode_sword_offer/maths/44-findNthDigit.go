package maths

import "strconv"

// @Author XuPEngHao
// @DATE 2022/6/18 19:48

/*
剑指 Offer 44. 数字序列中某一位的数字
数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

请写一个函数，求任意第n位对应的数字。



示例 1：

输入：n = 3
输出：3
示例 2：

输入：n = 11
输出：0


限制：


*/
// 123456789 - 1*1*9
// 10...99  - 2*10*9
// 100..999 ，3*100*9
// 1000..9999 ，4*1000*9
// 范围= 【位数 * start * 9 】从而就能计算出 n的处于的位置

// 0-9的= 10
// 10-99= 180
// 100-109,

// 输入的n  和数组的长度有什么关系吗？
// 15-2
// 14-1
// 13-1
// 12-1
// 11-0
// 10-1
func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count // 减去 1位，2位的count
		digit += 1
		start *= 10
		count = digit * start * 9 // 计算count的数量
	}

	// 当循环退出的时候，我们就能算出，开始位置start，位数digit

	// 计算单个
	var num = start + (n-1)/digit // 如何得出，这个计算公式，通过推理，找出n 和 start和n直接的关系
	numStr := strconv.FormatInt(int64(num), 10)
	return int(numStr[(n-1)%digit] - '0')
}
