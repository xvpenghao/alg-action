package bit

// @Author XuPEngHao
// @DATE 2022/5/28 13:52

/*
剑指 Offer 65. 不用加减乘除做加法
写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。



示例:

输入: a = 1, b = 1
输出: 2


提示：

a, b 均可能是负数或 0
结果不会溢出 32 位整数
*/

// https://leetcode.cn/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/solution/dian-zan-yao-wo-zhi-dao-ni-xiang-kan-dia-ovxy/
func add(a int, b int) int {
	// 思考10进制的操作，在思考二进制的操作，进位和本位
	var carry int
	for b != 0 {
		carry = (a & b) << 1 // 计算进位
		a = a ^ b            // 计算本位
		b = carry
	}
	return a
}
