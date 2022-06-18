package recursion

// @Author XuPEngHao
// @DATE 2022/6/18 19:06

/*
剑指 Offer 64. 求1+2+…+n
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。



示例 1：

输入: n = 3
输出: 6
示例 2：

输入: n = 9
输出: 45

*/

func sumNumsV2(n int) int {
	var res int
	var sumFn func(n int) bool
	sumFn = func(n int) bool {
		res += n
		return n > 0 && sumFn(n-1)
	}
	sumFn(n)
	return res
}

var res int

func sumNums(n int) int {
	_ = n > 1 && sumNums(n-1) > 0
	res += n
	return res
}

func sumNums2(n int) int {
	if n == 1 { // 不用if
		return 1
	}
	return n + sumNums(n-1)
}
