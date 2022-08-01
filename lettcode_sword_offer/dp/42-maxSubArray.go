package dp

import "math"

// @Author XuPEngHao
// @DATE 2022/7/6 22:34

/*
剑指 Offer 42. 连续子数组的最大和
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。



示例1:

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。


提示：

1 <= arr.length <= 10^5
-100 <= arr[i] <= 100
*/

/*
动态规划 (dp)
状态定义：dp[i]是代表以nums[i]为结尾的连续子数组的最大和。
转义方程：db的状态转义
dp[i-1] >0 则 db[i] = db[i-1]+  nums[i]
db[i-1]<=0 db[i] = nums[i],
这样判断是 避免 dp[i-1] 可能会对结果值，操作负的影响。

初值状态：db[0] = nums[0]
返回值：db数组的中的最大值
*/

func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		// 利用max，巧妙的化解了if else的判断逻辑，
		nums[i] += int(math.Max(float64(nums[i-1]), 0))
		res = int(math.Max(float64(nums[i]), float64(res)))
	}
	return res
}
