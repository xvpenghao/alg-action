package sword_offer

/**
* @Author XuPEngHao
* @Date 2023/7/2 16:55
 */

/*
剑指 Offer 20. 表示数值的字符串
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。

数值（按顺序）可以分成以下几个部分：

若干空格
一个 小数 或者 整数
（可选）一个 'e' 或 'E' ，后面跟着一个 整数
若干空格
小数（按顺序）可以分成以下几个部分：

（可选）一个符号字符（'+' 或 '-'）
下述格式之一：
至少一位数字，后面跟着一个点 '.'
至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
一个点 '.' ，后面跟着至少一位数字
整数（按顺序）可以分成以下几个部分：

（可选）一个符号字符（'+' 或 '-'）
至少一位数字
部分数值列举如下：

["+100", "5e2", "-123", "3.1416", "-1E-16", "0123"]
部分非数值列举如下：

["12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"]
*/
func isNumber(s string) bool { // 有限状态自动机器
	// 0-8中的状态集合，根据状态转移图，来设置状态map
	// 数值，空格+小或者整数+eE+空格
	mapList := []map[string]int{
		{" ": 0, "s": 1, "d": 2, ".": 4}, // 起始的空格，
		{"d": 2, ".": 4},                 // e的之前的sign +-
		{"d": 2, ".": 3, "e": 5, " ": 8},
		{"d": 3, "e": 5, " ": 8},
		{"d": 3},
		{"s": 6, "d": 7},
		{"d": 7},
		{"d": 7, " ": 8},
		{" ": 8},
	}
	p := 0
	t := ""
	for _, s1 := range s {
		ss1 := string(s1)
		// 字符串的 s1转换
		switch {
		case ss1 >= "0" && ss1 <= "9":
			t = "d" // 表示数字
		case ss1 == "+" || ss1 == "-":
			t = "s" // 表示符号
		case ss1 == "e" || ss1 == "E":
			t = "e" // 表示e
		case ss1 == "." || ss1 == " ":
			t = ss1 // 表示数字
		default:
			t = "?" // 其他
		}
		if _, ok := mapList[p][t]; !ok {
			return false
		}
		p = mapList[p][t] // 获取新的状态转移下标
	}
	// 这个 2，3,7,8的由来，可以根据状态机啦和题目来确定 数值（按顺序）可以分成以下几个部分：
	// 只有2，3,7，8是到 最后的空格的
	return p == 2 || p == 3 || p == 7 || p == 8
}
