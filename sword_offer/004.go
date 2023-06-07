package sword_offer

/**
* @Author XuPEngHao
* @Date 2023/6/6 21:18
 */

/*
剑指 Offer 04. 二维数组中的查找
在一个 n * m 的二维数组中，每一行都按照从左到右 非递减 的顺序排序，每一列都按照从上到下 非递减 的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。



示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false
限制：

0 <= n <= 1000

0 <= m <= 1000
*/

// 顺序有序，则可以相等 有序的二叉树
// i行特性 i尾最大
// j列特性，j首最小
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	// 左右两边没有更多数字了，则代表结束
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] > target { // 消掉列 我最小的比你的还大，
			j--
		} else if matrix[i][j] < target { // 消掉列
			i++
		} else {
			return true
		}
	}
	return false
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, numsList := range matrix {
		if binarySearch(numsList, target) {
			return true
		}
	}
	return false
}

// 二分查找,一行一行的
// [2,   5,  8, 12, 19] 单调增
func binarySearch(numList []int, target int) bool {
	start, end := 0, len(numList)-1
	for start <= end { // 参考target来移动的
		mid := (start + end) >> 1
		if numList[mid] > target { // 说明target再左边
			end = mid - 1
		} else if numList[mid] < target { // 说明target再右边,搜西
			start = mid + 1
		} else if numList[mid] == target { // ==
			return true
		}
	}
	return false
}
