package hash_table

import (
	"strings"
)

// @Author XuPEngHao
// @DATE 2022/4/20 09:41

/*
剑指 Offer 48. 最长不含重复字符的子字符串
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。



示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


提示：

s.length <= 40000
*/
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	maxStr := string(s[0])
	bak := []byte{s[0]}
	index := -1
	for i := 1; i < len(s); i++ {
		index = strings.LastIndex(string(bak), string(s[i]))
		if index != -1 {
			bak = bak[index+1:]
		}
		bak = append(bak, s[i])
		if len(bak) > len(maxStr) {
			maxStr = string(bak)
		}
	}
	return len(maxStr)
}
