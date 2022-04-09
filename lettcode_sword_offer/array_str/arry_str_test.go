package array_str

import (
	"fmt"
	"testing"
)

// @Author XuPEngHao
// @DATE 2022/3/12 11:30

func Test_strToInt(t *testing.T) {
	t.Log(strToIntV2("   -42"))
	t.Log(strToIntV2("+-42"))
	t.Log(strToIntV2("  +42word 34"))
	t.Log(strToIntV2("-2147483647"))
	t.Log(strToInt3("-2147483647"))

}

func Test_constructArr(t *testing.T) {
	t.Log(constructArr2([]int{1, 2, 3, 4, 5}))
	t.Log(constructArr([]int{1, 2, 3, 4, 5}))
}

func Test_isStraight(t *testing.T) {
	t.Log(isStraightV2([]int{0, 0, 2, 2, 5}))
	t.Log(isStraight([]int{0, 0, 1, 2, 5}))
	t.Log(isStraight([]int{1, 2, 0, 4, 5}))
	t.Log(isStraight([]int{0, 0, 2, 2, 5}))
	t.Log(isStraight([]int{11, 8, 12, 8, 10}))
}

func Test_reverseLeftWords(t *testing.T) {
	t.Log(reverseLeftWords("abcdefg", 2))
	t.Log(reverseLeftWords2("abcdefg", 2))
	t.Log(reverseLeftWords3("abcdefg", 2))
}

func Test_reverseWords(t *testing.T) {
	t.Log(reverseWords("  hello world! "))
}

func Test_getInsertLoc(t *testing.T) {
	t.Log(twoSum3([]int{10, 26, 30, 31, 47, 60}, 40))
	t.Log(twoSum3([]int{2, 7, 11, 15}, 9))
	t.Log(getInsertLoc([]int{2, 7, 11, 15}, 11))
	t.Log(binarySearch3([]int{2, 7, 11, 15}, 9))
}

func Test_findContinuousSequence(t *testing.T) {
	t.Log(findContinuousSequence(15))
	// t.Log(findContinuousSequence2(15))
	t.Log(findContinuousSequence3(15))
	// t.Log(findContinuousSequence(98160))
}

func Test_missingNumber(t *testing.T) {
	t.Log(missingNumber([]int{0, 1, 3, 4, 5}))
}

func Test_search(t *testing.T) {
	// t.Log(binarySearch2([]int{2, 2}, 3))
	t.Log(searchV3([]int{-109, -109, -109, 5, 7, 7, 8, 8, 10}, -109))
}

func Test_majorityElement(t *testing.T) {
	t.Log(majorityElementV4([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
	t.Log(majorityElementV4([]int{8, 8, 7, 7, 7}))
}

func Test_spiralOrder(t *testing.T) {
	/*t.Log(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))*/
	// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
	t.Log(spiralOrderV2([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
}

func Test_minNumber(t *testing.T) {
	// 3,30,34,5,9
	t.Log(minNumber([]int{20, 1}))
	t.Log(minNumber([]int{10, 1}))
}

func Test_exchange(t *testing.T) {
	t.Log(4 & 1) // 100
	t.Log(3 & 1) // 101
	fmt.Println(exchangeV3([]int{1, 3, 5}))
}
func Test_printNumbers(t *testing.T) {
	t.Log(printNumbers(2))
}

func Test_minArray(t *testing.T) {
	// t.Log(minArray([]int{2, 2, 2, 0, 1}))
	// t.Log(minArrayV2([]int{2, 2, 2, 0, 1}))
	// t.Log(minArrayV2([]int{3, 4, 5, 1, 2}))
	// t.Log(minArrayV3([]int{6, 1, 2, 3, 4, 5}))
	t.Log(minArrayV3([]int{3, 3, 1, 3}))
}

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
