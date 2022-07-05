package dp

// @Author XuPEngHao
// @DATE 2022/6/18 21:52

/*
剑指 Offer 10- I. 斐波那契数列
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。



示例 1：

输入：n = 2
输出：1
示例 2：

输入：n = 5
输出：5

*/
// 0 1 2 3 5 8 13
func fibV2(n int) int {
	a, b := 0, 1
	var sum int
	for i := 0; i < n; i++ {
		sum = (a + b) % 1000000007
		a = b // sum其实是多算了一天的，a=b（前一天算的）
		b = sum
	}
	return a
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
