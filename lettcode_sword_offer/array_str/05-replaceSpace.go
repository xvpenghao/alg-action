package array_str

import (
	"bytes"
	"strings"
)

// @Author XuPEngHao
// @DATE 2022/3/15 08:15

/*
剑指 Offer 05. 替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。



示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."
*/

func replaceSpaceV2(s string) string {
	if s == "" {
		return ""
	}
	// 计算出出空格个数
	spaceCnt := len(strings.Split(s, " ")) - 1
	// 字符串的数量 - 空格的数量 + （空格数 * len(%20)的长度）
	resLen := len(s) - spaceCnt + (spaceCnt * len("%20"))
	resList := make([]byte, 0, resLen)
	for index := range s {
		if s[index] != ' ' {
			resList = append(resList, s[index])
			continue
		}
		resList = append(resList, '%', '2', '0')
	}
	return string(resList)
}

func replaceSpace(s string) string {
	if s == "" {
		return ""
	}
	splitLen := len(strings.Split(s, " "))
	resList := make([]string, splitLen+splitLen-1)
	bb, i := bytes.Buffer{}, 0
	for index, val := range s {
		if val != ' ' {
			bb.WriteByte(s[index])
			continue
		}
		resList[i] = bb.String()
		i++
		resList[i] = "%20"
		i++
		bb.Reset()
	}
	// 到最后的 bb 可能还有值
	if bb.Len() > 0 {
		resList[i] = bb.String()
		bb.Reset() // 保持好习惯
	}
	return strings.Join(resList, "")
}
