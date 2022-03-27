package array_str

// @Author XuPEngHao
// @DATE 2022/3/22 08:40

/*
剑指 Offer 29. 顺时针打印矩阵
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。



示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：


输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]



限制：

0 <= matrix.length <= 100
0 <= matrix[i].length <= 100
*/

/*
1，2，3，
4，5，6，
7，8，9，
*/

/*
找规律，做图分析
*/
func spiralOrderV2(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	result := make([]int, len(matrix)*len(matrix[0])) // 提前分配好指定空间
	i := 0
	for {
		// 左-右,top不动
		for j := left; j <= right; j++ {
			result[i] = matrix[top][j]
			i++
		}
		if top++; top > bottom {
			break
		}

		// 上-下，right不动
		for j := top; j <= bottom; j++ {
			result[i] = matrix[j][right]
			i++
		}
		if right--; left > right {
			break
		}

		// 右-左，bottom不动
		for j := right; j >= left; j-- {
			result[i] = matrix[bottom][j]
			i++
		}
		if bottom--; top > bottom {
			break
		}

		// 下-上，left不动
		for j := bottom; j >= top; j-- {
			result[i] = matrix[j][left]
			i++
		}
		if left++; left > right {
			break
		}
	}

	return result
}

/*
rows,cols，做好边界处理
结束的条件：len(res)==len(matrix)的长度
左，下，右，上，
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	rows, cols := len(matrix)-1, len(matrix[0])-1
	row, col, i, eleTotalCnt := 0, 0, 0, len(matrix)*len(matrix[0])
	result := make([]int, eleTotalCnt) // 提前分配好指定空间
	for row <= rows && col <= cols && i < eleTotalCnt {
		// 左, row不动，col++，执行完，row++（丢弃）
		for j := col; j <= cols; j++ {
			result[i] = matrix[row][j]
			i++
		}
		row++

		// 下- cols-1不动，row++，执行完，cols--(丢弃)
		for j := row; j <= rows && i < eleTotalCnt; j++ {
			result[i] = matrix[j][cols]
			i++
		}
		cols--

		// 右, rows-1不动，cols--，执行完，rows--（丢弃）
		for j := cols; j >= col && i < eleTotalCnt; j-- {
			result[i] = matrix[rows][j]
			i++
		}
		rows--

		// 上，col不动，rows--，执行完毕，col++（丢弃）
		for j := rows; j >= row && i < eleTotalCnt; j-- {
			result[i] = matrix[j][col]
			i++
		}
		col++
	}

	return result
}
