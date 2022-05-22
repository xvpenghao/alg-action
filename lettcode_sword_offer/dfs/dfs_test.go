package dfs

import "testing"

// @Author XuPEngHao
// @DATE 2022/5/7 08:46

func Test_exist(t *testing.T) {
	board := [][]byte{{'a', 'b'}, {'c', 'c'}}
	t.Log(exist(board, "ab"))
}
