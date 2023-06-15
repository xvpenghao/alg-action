package alg_sort

import "testing"

/**
* @Author XuPEngHao
* @Date 2023/6/15 07:27
 */

func Test_bubble(t *testing.T) {
	// 快速的排序
	//numList := []int{1, 5, 3}
	numList := []int{1, 2, 3, 4}
	quickSortV2(numList, 0, len(numList)-1)
	t.Log(numList)
}
