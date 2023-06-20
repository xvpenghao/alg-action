package sword_offer

import "bytes"

/**
* @Author XuPEngHao
* @Date 2023/6/20 08:26
 */

/*
剑指 Offer 05. 替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。



示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."


限制：

0 <= s 的长度 <= 10000
*/

func replaceSpace(s string) string {
	buf := bytes.Buffer{}
	for _, val := range s {
		if val == ' ' {
			buf.WriteString("%20")
		} else {
			buf.WriteString(string(val))
		}
	}
	return buf.String()
}
