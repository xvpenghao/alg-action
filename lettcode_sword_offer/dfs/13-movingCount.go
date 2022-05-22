package dfs

// @Author XuPEngHao
// @DATE 2022/5/22 10:54

/*
剑指 Offer 13. 机器人的运动范围
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？



示例 1：

输入：m = 2, n = 3, k = 1
输出：3
示例 2：
00，01，02
10，11，12

输入：m = 3, n = 1, k = 0
输出：1
提示：

1 <= n,m <= 100
0 <= k <= 20
*/

// 这块我理解错误了一个地方，我以上以为好走2次呢，
// 第一次-先左走，第二次在下走，两次的起始位置都是从00开始，这块我还是对dfs理解不太深入
// dfs,深度优先遍历，一条路走到嘿，不通，然后撤出，然后在接着好下一个出路。

func movingCount(m int, n int, k int) int {

	// 构造矩阵，m*n
	// 起始点是固定的，先左/先下计算 累计计算走的个数
	// 计算 行列位数相加是否满足的k的要求
	var board [][]int
	for i := 0; i < m; i++ {
		board = append(board, make([]int, n))
	}
	return dfsMC(board, 0, 0, k)
}

func dfsMC(board [][]int, i, j, k int) int {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || board[i][j] != 0 || !isCanFront(i, j, k) {
		return 0
	}
	board[i][j] = -1 // 用户布尔值可以的
	return 1 + dfsMC(board, i, j+1, k) + dfsMC(board, i+1, j, k) + dfsMC(board, i-1, j, k) + dfsMC(board, i, j-1, k)
}

func movingCount2(m int, n int, k int) int {
	var board [][]int
	for i := 0; i < m; i++ {
		board = append(board, make([]int, n))
	}
	dfsMC2(board, 0, 0, k)
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt += board[i][j]
		}
	}
	return cnt
}

func dfsMC2(board [][]int, i, j, k int) {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || board[i][j] != 0 || !isCanFront(i, j, k) {
		return
	}
	board[i][j] = 1
	dfsMC2(board, i, j+1, k)
	dfsMC2(board, i+1, j, k)
	dfsMC2(board, i-1, j, k)
	dfsMC2(board, i, j-1, k)
}

func isCanFront(i, j, k int) bool {
	sum := 0
	for i > 0 || j > 0 {
		sum += i % 10
		i /= 10
		sum += j % 10
		j /= 10
	}
	return sum <= k
}
