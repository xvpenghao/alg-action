package hash_table

import "testing"

// @Author XuPEngHao
// @DATE 2022/4/19 09:01

func Test_firstUniqChar(t *testing.T) {
	t.Log(string(firstUniqChar("dddccdbba")))
	t.Log(string(firstUniqCharV2("dddccdbba")))
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcabcbb"))
	t.Log(lengthOfLongestSubstring("bbbbb"))
	t.Log(lengthOfLongestSubstring("pwwkew"))
}

func Test_findRepeatNumber(t *testing.T) {
	// t.Log(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
	t.Log(findRepeatNumberV3([]int{2, 3, 1, 0, 2, 5, 3}))
	// t.Log(findRepeatNumberV3([]int{1, 2, 3, 4, 5, 5}))
}
