package array_str

import (
	"bytes"
	"sort"
	"strconv"
)

// @Author XuPEngHao
// @DATE 2022/3/21 19:44
/*
剑指 Offer 45. 把数组排成最小的数
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。



示例 1:

输入: [10,2]
输出: "102"
示例 2:

输入: [3,30,34,5,9]
输出: "3033459"


提示:

0 < nums.length <= 100
说明:

输出结果可能非常大，所以你需要返回一个字符串而不是整数
拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
*/
// [3,30]
// x='3',y='30'
// x+y > y+x, 3 30 > 30 3 则说明  x >y ,则 将 y 与x交换位置即可
// 问题的核心在于，找出排序的规则。
func minNumber(nums []int) string {
	// 自定义排序规则
	sort.SliceStable(nums, func(i, j int) bool {
		x := strconv.Itoa(nums[i])
		y := strconv.Itoa(nums[j])
		// x = 20
		// y = 1
		return x+y < y+x
	})
	b := bytes.Buffer{}
	for _, val := range nums {
		b.WriteString(strconv.Itoa(val))
	}
	return b.String()
}
