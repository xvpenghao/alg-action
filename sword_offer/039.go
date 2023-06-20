package sword_offer

/**
* @Author XuPEngHao
* @Date 2023/6/19 08:43
 */
/*
剑指 Offer 39. 数组中出现次数超过一半的数字
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。


示例 1:

输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
[1, 2, 3, 2, 2, 2, 5, 4, 2]
[1, 2, 2, 2, 2, 2, 3, 4, 5]
输出: 2
*/

func majorityElement(nums []int) int {
	// 以为本身的占比大，所以丢了几个不是问题
	repeatMaxNum := nums[0]
	repeatCnt := 0
	for _, val := range nums {
		if repeatCnt == 0 {
			repeatMaxNum = val
		}
		if repeatMaxNum == val {
			repeatCnt++
		} else {
			repeatCnt--
		}
	}
	return repeatMaxNum
}
