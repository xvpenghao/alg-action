package array_str

import "bytes"

// @Author XuPEngHao
// @DATE 2022/4/5 12:09
/*
剑指 Offer 58 - II. 左旋转字符串
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。



示例 1：

输入: s = "abcdefg", k = 2
输出: "cdefgab"
示例 2：

输入: s = "lrloseumgh", k = 6
输出: "umghlrlose"


限制：

1 <= k < s.length <= 10000
*/

// 使用求余数法则
func reverseLeftWords3(s string, n int) string {
	buf := bytes.Buffer{}
	end := n + len(s)
	for i := n; i < end; i++ {
		buf.WriteByte(s[i%len(s)]) // 秒呀
	}
	return buf.String()
}

// 遍历
func reverseLeftWords2(s string, n int) string {
	buf := bytes.Buffer{}
	end := len(s)
	for i := n; i < end; i++ {
		buf.WriteByte(s[i])
		if len(s) == i+1 {
			i, end = -1, n
		}
	}
	return buf.String()
}

func reverseLeftWords(s string, n int) string {
	if s == "" {
		return ""
	}
	return s[n:] + s[:n]
}
