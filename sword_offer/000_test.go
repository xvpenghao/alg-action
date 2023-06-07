package sword_offer

import "testing"

/**
* @Author XuPEngHao
* @Date 2023/6/5 11:04
 */

func Test_011(t *testing.T) {
	numbers := []int{3, 4, 5, 1, 2}
	numbers = []int{2, 2, 2, 0, 1}
	numbers = []int{1, 3, 5}
	numbers = []int{3, 1, 3, 3}
	t.Log(minArray(numbers))
}

func Test_007(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)
	t.Log(root)
}

func Test_004(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	t.Log(findNumberIn2DArray2(matrix, 18))
}

func Test_003(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	nums = []int{3, 4, 2, 1, 1, 0}
	t.Log(findRepeatNumberV3(nums))
}
