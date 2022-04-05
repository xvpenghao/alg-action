package array_str

import (
	"bytes"
	"strings"
)

// @Author XuPEngHao
// @DATE 2022/4/4 20:24

/*
剑指 Offer 58 - I. 翻转单词顺序
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。


示例 1：

输入: "the sky is blue"
输出:"blue is sky the"
示例 2：

输入: "  hello world! "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good  example"
输出:"example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。


说明：

无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
*/

// 双指针
func reverseWords2(s string) string {
	buf := bytes.Buffer{}
	start, end := len(s)-1, len(s)-1 // start 本来就是下标
	for start >= 0 {
		// 找到不为空格的
		for start >= 0 && s[start] != ' ' {
			start--
		}
		// 末尾可能会添加的多于的空格，对于一上来末尾就是空格的的来说
		buf.WriteString(s[start+1:end+1] + " ")
		// 过滤空格
		for start >= 0 && s[start] == ' ' {
			start--
		} // start本来是空格的位置的下标，现在-1了

		// 空格过滤完毕后则，end和start是一致的，
		end = start
	}
	// 可能最后会多加空格
	return strings.TrimSpace(buf.String())
}

// 注意事项，
// 反转后的单词 首部和尾部不能包含空格，
// 反转后的单词之间的空格只能有个一个。
func reverseWords(s string) string {
	buf := bytes.Buffer{}
	strList := strings.Split(s, " ")
	for i := len(strList) - 1; i >= 0; i-- {
		if strList[i] == "" { // 找到分隔符，然后写入值,连续出现空格的次数为则写入改空格
			continue
		}
		buf.WriteString(" ")
		buf.WriteString(strList[i])
	}
	return strings.TrimSpace(buf.String())
}
