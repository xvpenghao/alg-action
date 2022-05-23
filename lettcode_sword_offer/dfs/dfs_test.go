package dfs

import (
	"fmt"
	"testing"
)

// @Author XuPEngHao
// @DATE 2022/5/7 08:46

func Test_mirrorTree(t *testing.T) {
	res := mirrorTreeV2(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	})
	t.Log(res)
}

func Test_movingCount(t *testing.T) {
	fmt.Println(movingCount(2, 3, 1))
	fmt.Println(movingCount(3, 1, 0))
	t.Log("movingCount2")
	fmt.Println(movingCount2(2, 3, 1))
	fmt.Println(movingCount2(3, 1, 0))
}

func Test_exist(t *testing.T) {
	board := [][]byte{{'a', 'b'}, {'c', 'c'}}
	t.Log(exist(board, "ab"))
}
