package bit

// @Author XuPEngHao
// @DATE 2022/5/28 09:43

/*
剑指 Offer 56 - I. 数组中数字出现的次数
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。


示例 1：

输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]
示例 2：

输入：nums = [1,2,10,4,1,4,3,3]
输出：[2,10] 或 [10,2]


限制：

2 <= nums.length <= 10000
*/

func singleNumbersV2(nums []int) []int {
	x, y, n, m := 0, 0, 0, 1
	for _, num := range nums {
		n ^= num
	}
	// 将nums分层两组尽心异或，如何分组呢，最简单的就是，奇偶分组，也就是&1操作
	// 相同的数可定会分到一组，如何将不相同的连个数，分配到2个组中
	// num1 ^ num2 的结果
	// 将该结果进行拆分
	// 异或结果 111 ^ 001 = 110
	// 110如何划分成 2组呢？，核心就找出 这两个数的mask吗记可以了，最地位首个位1的
	for n&m == 0 { // 异或的结果值，然后 & 操作找到首位为1的，则将nums划分为2组
		m = m << 1 // 往左移*2
	}
	for _, num := range nums {
		if num&m > 0 {
			x ^= num
		} else {
			y ^= num
		}
	}
	return []int{x, y}
}

// [2,1,3,2,3]
func singleNumbers(nums []int) []int {
	m := make(map[int]int, len(nums))
	for _, n := range nums {
		m[n] = m[n] + 1
	}
	result := make([]int, 0, 2)
	for n, cnt := range m {
		if cnt > 1 {
			continue
		}
		result = append(result, n)
	}
	return result
}
