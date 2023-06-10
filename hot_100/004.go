package hot_100

import (
	"sort"
)

/**
* @Author XuPEngHao
* @Date 2023/6/9 14:08
 */

/*
4. 寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。



示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

从小到大
*/

// 【1，,3】 [2]
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	// 1、合并数组的长度我知道后，我就得出，中位数的下标是那几个
	sumLen := len(nums1) + len(nums2)
	m1Index := sumLen / 2
	m2Index := -1
	if sumLen%2 == 0 {
		m2Index = m1Index - 1
	}

	n1Idx, n2Idx := 0, 0
	m1Num, m2Num := 0, 0
	for i := 0; i < sumLen; i++ {
		b1 := false
		b2 := false
		if n1Idx < len(nums1) && n2Idx < len(nums2) {
			if nums1[n1Idx] < nums2[n2Idx] {
				n1Idx++
				b1 = true
			} else {
				n2Idx++
				b2 = true
			}
		} else {
			if n1Idx < len(nums1) {
				n1Idx++
				b1 = true
			}
			if n2Idx < len(nums2) {
				n2Idx++
				b2 = true
			}
		}

		// 【1，,3】 [2]
		if i == m1Index || i == m2Index {
			if b1 {
				m1Num += nums1[n1Idx-1]
			}
			if b2 {
				m2Num += nums2[n2Idx-1]
			}
		}
	}

	if m2Index == -1 {
		return float64(m1Num) + float64(m2Num)
	}
	return (float64(m1Num) + float64(m2Num)) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 暴力的解法
	nums1 = append(nums1, nums2...)
	sort.SliceStable(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})

	n1Index := len(nums1) / 2
	if len(nums1)%2 != 0 {
		return float64(nums1[n1Index])
	}
	return (float64(nums1[n1Index]) + float64(nums1[n1Index-1])) / 2
}
