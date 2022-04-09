package array_str

// @Author XuPEngHao
// @DATE 2022/4/8 10:00

/*
剑指 Offer 66. 构建乘积数组
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。



示例:

输入: [1,2,3,4,5]
输出: [120,60,40,30,24]

[1,2,3,4,5]
[120,60,40,30,24]
提示：

所有元素乘积之和不会溢出 32 位整数
a.length <= 100000

*/

func constructArr3(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := make([]int, len(nums))
	// 存放 计算左边累计乘积
	result[0] = 1 // L [1 1 2 6 24]
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	right := 1 // 动态的来构造 右边数据
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * right
		right *= nums[i]
	}
	return result
}

// 这个数的左边* 这个数右边，核心问题就是，计算好这个数的，左边和右边的问题
func constructArr2(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	L, R, result := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	// 存放 计算左边累计乘积
	L[0] = 1 // L [1 1 2 6 24]
	for i := 1; i < len(nums); i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	// 计算右边
	R[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}

	// 然后乘积
	for i := 0; i < len(nums); i++ {
		result[i] = L[i] * R[i]
	}
	return result
}

// 超出时间限制
func constructArr(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = 1
		for j := 0; j < len(nums); j++ {
			if i != j {
				result[i] *= nums[j]
			}
		}
	}
	return result
}
