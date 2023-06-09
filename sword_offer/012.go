package sword_offer

/**
* @Author XuPEngHao
* @Date 2023/6/8 09:00
 */

/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

 

例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。



 

示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false
 

提示：

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board 和 word 仅由大小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 已经用过的不就不能再用啦
// 题中值给出了，word存在于二维数组，单丝并未说明，他是在开始还是结束，所以需要行遍历来找寻
func exist(board [][]byte, word string) bool {
	iend := len(board) - 1
	jend := len(board[0]) - 1
	for i, nums := range board {
		for j := range nums {
			if dfs(board, i, j, iend, jend, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, iend, jend int, word string, wordIndex int) bool {
	// 下标越界或者 不等于 words则返回 false
	if i < 0 || i > iend || j < 0 || j > jend || board[i][j] != word[wordIndex] {
		return false
	}

	if wordIndex == len(word)-1 { // 表示全部对比完毕
		return true
	}

	// 标记已经被使用过啦
	board[i][j] = '0'
	bol := dfs(board, i+1, j, iend, jend, word, wordIndex+1) || dfs(board, i-1, j, iend, jend, word, wordIndex+1) ||
		dfs(board, i, j+1, iend, jend, word, wordIndex+1) || dfs(board, i, j-1, iend, jend, word, wordIndex+1)
	board[i][j] = word[wordIndex]
	return bol
}
