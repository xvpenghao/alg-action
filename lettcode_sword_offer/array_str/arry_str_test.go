package array_str

import (
	"testing"
)

// @Author XuPEngHao
// @DATE 2022/3/12 11:30

func Test_replaceSpace(t *testing.T) {
	t.Log(replaceSpaceV2("a b c"))
	// numList := make([]int, 0, 2) 好处是什么
}

func Test_binarySearch(t *testing.T) {
	numList := []int{1, 3, 6, 10, 25, 30}
	t.Log(binarySearch(numList, 30))
}

func Test_findNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		// {10, 13, 14, 17, 24},
		// {18, 21, 23, 26, 30},
	}
	// t.Log(matrix)
	t.Log(findNumberIn2DArrayV3(matrix, 100))
}

func TestName(t *testing.T) {
	numList := [][]int{{}, {1}}
	t.Log(len(numList), len(numList[0]), len(numList[1]))
}
