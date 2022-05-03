package hash_table

// @Author XuPEngHao
// @DATE 2022/4/28 18:00

/*
剑指 Offer 50. 第一个只出现一次的字符
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

fist =
patch
s = dddccdbba


示例 1:

abcwcabdc

输入：s = "abaccdeff"
输出：'b'
示例 2:

输入：s = ""
输出：' '


限制：

0 <= s 的长度 <= 50000
*/

func firstUniqCharV2(s string) byte {
	if s == "" {
		return ' '
	}
	m := make(map[byte]bool, len(s))
	uniqStrList := make([]byte, 0)
	for _, val := range s {
		_, isFirst := m[byte(val)]
		m[byte(val)] = !isFirst // 如果是第一次，则肯定是false，然后则取反
		if !isFirst {
			uniqStrList = append(uniqStrList, byte(val))
		}
	}

	for _, val := range uniqStrList {
		if bol, _ := m[val]; bol {
			return val
		}
	}
	return ' '
}

func firstUniqChar(s string) byte {
	if s == "" {
		return ' '
	}
	m := make(map[byte]bool, len(s))
	for _, val := range s {
		_, isFirst := m[byte(val)]
		m[byte(val)] = !isFirst // 如果是第一次，则肯定是false，然后则取反
	}

	for i := 0; i < len(s); i++ {
		if val, _ := m[s[i]]; val {
			return s[i]
		}
	}
	return ' '
}
