package array_str

import (
	"math"
	"sort"
)

// @Author XuPEngHao
// @DATE 2022/4/5 19:37

/*
剑指 Offer 61. 扑克牌中的顺子
从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。



示例 1:

输入: [1,2,3,4,5]
输出: True


示例 2:

输入: [0,0,1,2,5]
输出: True


限制：

数组长度为 5

数组的数取值为 [0, 13] .
*/

func isStraightV3(nums []int) bool {
	// 里面大小王，则排到前面
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	joke := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 { // 0的直接忽略，
			joke++
		} else if nums[i] == nums[i+1] { // 如果里面有相等的，直接，return false
			return false
		}
	}
	return nums[4]-nums[joke] < 5 // max-min 锁定范围
}

// 5张牌，是顺子，则利用的 [min,max]就可以范围进行锁定了，
// 如果真的的是 顺子，则max -min <5 利用这个特性，去思考问题
func isStraightV2(nums []int) bool {
	var max float64 = 0
	var min float64 = 14
	// 如果有重复的数据则直接返回false，0忽略
	set := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 { // 0的直接忽略，
			continue
		}
		// 非0的在判断是否有重复数据
		if _, ok := set[nums[i]]; ok { // 有重复数据则直接返回false
			return false
		} else {
			set[nums[i]] = struct{}{}
		}

		max = math.Max(max, float64(nums[i]))
		min = math.Min(min, float64(nums[i]))
	}
	return max-min < 5 // max-min 锁定范围
}

func isStraight(nums []int) bool {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	zeroCtn := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			zeroCtn++
			continue
		}
		if nums[i+1]-nums[i] == 1 { // 前后错1，则不需要补充
			continue
		}
		if nums[i+1]-nums[i] == 0 || nums[i+1]-nums[i]-1 > 2 { // 有相等或者，错的》2的直接退出
			return false
		}
		// 这是你要补的
		zeroCtn -= nums[i+1] - nums[i] - 1
		if zeroCtn < 0 {
			return false
		}
	}
	return true
}
