package sword_offer

/**
* @Author XuPEngHao
* @Date 2023/7/2 18:23
 */

/*
剑指 Offer 38. 字符串的排列
输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
示例:

输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]

限制：

1 <= s 的长度 <= 8
*/

// abc
// abc, acb,bac,bca,cab,cba
// 1234567
// 1234567,1324567,1,
// 全排列的计算公式

func permutation(s string) []string {
	var res []string

	ss := []byte(s)
	var dfsFn func(x int)
	dfsFn = func(x int) {
		if x == len(ss)-1 {
			res = append(res, string(ss))
			return
		}
		// 循环遍历
		setMap := make(map[byte]int, 0) // 避免添加重复测次
		for i := x; i < len(ss); i++ {
			if _, ok := setMap[ss[i]]; ok {
				continue
			}
			setMap[ss[i]] = 1
			ss[i], ss[x] = ss[x], ss[i]
			dfsFn(x + 1)
			ss[i], ss[x] = ss[x], ss[i] // 在恢复过啦
		}
	}

	dfsFn(0)
	return res
}

// 深度优先遍历
