package dfs

// @Author XuPEngHao
// @DATE 2022/5/7 08:46

/*
剑指 Offer 12. 矩阵中的路径
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。

示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false
*/
// 【如果网格中，存在word】，意味着，你搜索的的位置，可以是从任意位置开始。
// 为什么要使用dfs？
func exist(board [][]byte, word string) bool {
	if board == nil {
		return false
	}

	for i := 0; i < len(board); i++ { // 使用遍历矩阵操作，表示，words的词的开始位置，可能是，矩阵中的任意位置。
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) { // 每次都是从词的0个开始
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, i, j, k int) bool {
	// 判断是否满足条件
	if i > len(board)-1 || i < 0 || j > len(board[0])-1 || j < 0 || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 { // 词用完毕了，则返回true, 所以不需要考虑 k的下标越界的问题
		return true
	}
	board[i][j] = ' ' // 说明，改词可以被使用了，代表这个词我用过了
	// 因为都是或的关系，所以谁前前后，无所谓。开始查找下一个
	bol := dfs(board, word, i+1, j, k+1) || dfs(board, word, i-1, j, k+1) ||
		dfs(board, word, i, j+1, k+1) || dfs(board, word, i, j-1, k+1)
	board[i][j] = word[k] // 词的归还，便于下次，查找
	return bol
}
