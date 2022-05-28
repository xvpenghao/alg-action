package bit

// @Author XuPEngHao
// @DATE 2022/5/27 22:15

/*
剑指 Offer 16. 数值的整数次方
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。



示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25


提示：

-100.0 < x < 100.0
-231 <= n <= 231-1
-104 <= xn <= 104
*/
/*
3^10, 3
*/

// 快速幂，每次折半
func myPowV2(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	// 处理n为负数
	if n < 1 {
		x = 1 / x
		n = -n
	}
	// 3^10,
	// 10=9^5, 5=81^2*3, 2=81*81*3
	var result float64 = 1
	for n > 0 {
		if n&1 == 1 { // 是奇数，代表可能需要*一个x
			result *= x
		}
		x = x * x
		n = n >> 1
	}
	return result
}

// 超时
func myPow(x float64, n int) float64 {
	var result float64 = 1
	flags := false
	if n < 0 {
		n *= -1
		flags = true
	}
	for n > 0 {
		result *= x
		n--
	}
	if flags {
		result = 1 / result
	}
	return result
}
