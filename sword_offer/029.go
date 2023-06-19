package sword_offer

/**
* @Author XuPEngHao
* @Date 2023/6/15 08:57
 */
/*
剑指 Offer 29. 顺时针打印矩阵
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。



示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]
[1,2,3]
[4,5,6]
[7,8,9]
[1,2,3,6,9,8,7,4,5]
*/

// 什么时候我的循环才是结束,结束的条件是什么
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	if len(matrix) == 1 {
		return matrix[0]
	}

	result := make([]int, len(matrix)*len(matrix[0]))
	resIndex := 0
	left, top, right, below := 0, 0, len(matrix[0])-1, len(matrix)-1
	for true {
		// 从左到右，上+1
		for i := left; i <= right; i++ {
			result[resIndex] = matrix[top][i]
			resIndex++
		}
		if top = top + 1; top > below {
			break
		}
		// 从上到下，右-1
		for i := top; i <= below; i++ {
			result[resIndex] = matrix[i][right]
			resIndex++
		}
		if right = right - 1; right < left {
			break
		}

		// 从下到左，下-1
		for i := right; i >= left; i-- {
			result[resIndex] = matrix[below][i]
			resIndex++
		}
		if below = below - 1; below < top {
			break
		}

		// 下到上，左+1 你的等于没有问题是因为，在操作后，top的发生了变化
		for i := below; i >= top; i-- {
			result[resIndex] = matrix[i][left]
			resIndex++
		}
		if left = left + 1; left > right {
			break
		}
	}

	return result
}
