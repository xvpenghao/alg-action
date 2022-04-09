package array_str

import (
	"bytes"
	"math"
	"strings"
)

// @Author XuPEngHao
// @DATE 2022/4/9 17:08
/*
剑指 Offer 67. 把字符串转换成整数
写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。



首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。

当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。

注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。

在任何情况下，若函数不能进行有效的转换时，请返回 0。

说明：

假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。

示例 1:

输入: "42"
输出: 42
示例 2:

输入: "   -42"
输出: -42
解释: 第一个非空白字符为 '-', 它是一个负号。
     我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
示例 3:

输入: "4193 with words"
输出: 4193
解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
示例 4:

输入: "words and 987"
输出: 0
解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
     因此无法执行有效的转换。
示例 5:

输入: "-91283472332"
输出: -2147483648
解释: 数字 "-91283472332" 超过 32 位有符号整数范围。
     因此返回 INT_MIN (−231) 。
*/

// 丢弃无用的开头空格
// 忽略存在多余的字符
// 不能转换为返回0
/*
"  -0012a42"
输出：
-1242
预期结果：
-12
*/

// "1234"  res = 10*res+x // 每次多出个0 这样就+x
func strToIntV2(str string) int {
	str = strings.TrimSpace(str) // 如果不是函数，则去首部空格即可
	if str == "" {
		return 0
	}
	// binary 使用最后计算res是 需要判断的边界值
	min, max, binary := math.MinInt32, math.MaxInt32, math.MaxInt32/10
	sign := 1 // 符号
	i := 0    // 开始的位置
	if str[0] == '-' || str[0] == '+' {
		if str[0] == '-' {
			sign *= -1
		}
		i++
	}

	// 确定好，符号位置后，剩下的只要发现不是字符数字，则直接break，逻辑清晰

	var result int
	for ; i < len(str); i++ {
		if str[i] > '9' || str[i] < '0' { // 不是数字则break
			break
		}
		// 判断是否越级  max=最后一位7 ，
		if result > binary || (result == binary && str[i] > '7') {
			if sign == 1 {
				return max
			} else {
				return min
			}
		}
		result = 10*result + int(str[i]-'0')
	}

	return result * sign
}

// 错误的
func strToInt(str string) int {
	buf := bytes.Buffer{}
	sign2 := 1
	for i := 0; i < len(str); i++ {
		if buf.Len() == 0 {
			if str[i] == ' ' {
				continue
			}
			if str[i] == '+' || str[i] == '-' { // 是否符号，但是下一个不是数字则 return
				if i+1 < len(str) && (str[i+1] > '9' || str[i+1] < '0') {
					return 0
				}
			} else {
				if str[i] > '9' || str[i] < '0' {
					return 0
				}
			}
		}
		if str[i] == '-' {
			sign2 *= -1
			continue
		}
		if str[i] == '+' {
			continue
		}
		if str[i] > '9' || str[i] < '0' {
			break
		}
		if str[i] <= '9' && str[i] >= '0' {
			buf.WriteByte(str[i])
		}
	}

	if buf.Len() == 0 {
		return 0
	}

	bList := buf.Bytes()
	result := 0
	for i := 0; i < len(bList); i++ {
		result += int(bList[i]-'0') * int(math.Pow(10, float64(len(bList)-i-1)))
	}

	result = result * sign2
	if result < -1<<31 {
		result = -1 << 31
	}
	if result > math.MaxInt32 {
		result = math.MaxInt32
	}
	return result
}
