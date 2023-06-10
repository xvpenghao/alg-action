package hot_100

import "testing"

/**
* @Author XuPEngHao
* @Date 2023/6/9 10:57
 */

func Test_004(t *testing.T) {
	nums1 := []int{1, 4, 5}
	nums2 := []int{2, 3, 6}
	// 1,2,3,4,5,6
	t.Log(findMedianSortedArrays2(nums1, nums2))
}

func Test_001(t *testing.T) {
	numsList := []int{2, 7, 11, 15}
	target := 9

	numsList = []int{3, 2, 4}
	target = 6

	t.Log(twoSum2(numsList, target))
}
