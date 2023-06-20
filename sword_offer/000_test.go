package sword_offer

import (
	"fmt"
	"testing"
)

/**
* @Author XuPEngHao
* @Date 2023/6/5 11:04
 */

func Test_005(t *testing.T) {
	s := "We are happy."
	t.Log(replaceSpace(s))
}

func Test_039(t *testing.T) {
	numsList := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	t.Log(majorityElement(numsList))
}

func Test_031(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}

	//pushed = []int{1, 2, 3, 4, 5}
	//popped = []int{4, 3, 5, 1, 2}

	pushed = []int{1, 0, 2}
	popped = []int{2, 1, 0}

	t.Log(validateStackSequences(pushed, popped))
}

func Test_29(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	result := spiralOrder(matrix)
	t.Log(result)
}

// [1,1,0,0,1,0] 排序
// [1,1,1,0,0,0] 如何排序呢
func Test_021(t *testing.T) {
	numsList := []int{1, 1, 4, 8, 7, 20}
	numsList = []int{1, 1, 4, 8, 7, 20}
	//numsList = []int{1, 2, 3}
	t.Log(exchangeV2(numsList))
}

func Test_017(t *testing.T) {
	fmt.Println(printNumbers2(4))
}

func Test_012(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	t.Log(exist(board, "ABCCED"))
}

func Test_011(t *testing.T) {
	numbers := []int{3, 4, 5, 1, 2}
	numbers = []int{2, 2, 2, 0, 1}

	numbers = []int{3, 1, 3, 3}
	numbers = []int{1, 3, 5}
	t.Log(minArray2(numbers))
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
