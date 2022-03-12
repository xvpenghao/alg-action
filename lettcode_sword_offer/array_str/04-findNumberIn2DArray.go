package array_str

// @Author XuPEngHao
// @DATE 2022/3/11 09:48

/*
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target=5，返回true。

给定target=20，返回false。

限制：

0 <= n <= 1000

0 <= m <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 利用二维数组的，数组内容的特性
func findNumberIn2DArrayV3(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	// 从左下角开始计算
	row, col := len(matrix)-1, 0
	flg, colLen := 0, len(matrix[0])-1
	for row >= 0 && col <= colLen {
		flg = matrix[row][col]
		if target == flg {
			return true
		} else if target < flg { // 我是这一行最小的，如果 target比我好小，则它一定在我上方，消除该行
			row--
		} else if target > flg { //  我是这一列的最大的，如果 target比我大，则一定在我的右方。消除该列
			col++
		}
	}
	return false
}

func findNumberIn2DArrayV2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	rowLen, colLen := len(matrix), len(matrix[0])
	for i := 0; i < rowLen; i++ {
		if len(matrix[i]) == 0 {
			continue
		}
		// 对角线的不存在
		if !(i < colLen) {
			return false
		}
		if target >= matrix[i][i] && target <= matrix[i][colLen-1] && binarySearch(matrix[i], target) {
			return true
		}

		// 列的范围查询
		if target >= matrix[i][i] && target <= matrix[rowLen-1][i] {
			curRow, numList := i, make([]int, 0)
			for curRow <= rowLen-1 {
				numList = append(numList, matrix[curRow][i])
				curRow++
			}
			if binarySearch(numList, target) {
				return true
			}
		}
	}
	return false
}

// func xxx(matrix [][]int,)bool{
// 	start, end := 0, len(numList)-1
// 	for start <= end {
// 		mid := (start + end) / 2
// 		if numList[mid] == target {
// 			return true
// 		} else if numList[mid] > target {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// }

func binarySearch(numList []int, target int) bool {
	start, end := 0, len(numList)-1
	for start <= end {
		mid := (start + end) / 2
		if numList[mid] == target {
			return true
		} else if numList[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for i := range matrix {
		for j := range matrix[i] {
			if target == matrix[i][j] {
				return true
			}
		}
	}
	return false
}
