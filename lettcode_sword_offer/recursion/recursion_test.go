package recursion

import "testing"

// @Author XuPEngHao
// @DATE 2022/6/1 10:05

func Test_sumNums(t *testing.T) {
	t.Log(sumNums(2))
}

func Test_isSubStructure(t *testing.T) {
	A := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: -4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	B := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: -4,
		},
	}
	//
	// A = &TreeNode{
	// 	Val: 3,
	// 	Left: &TreeNode{
	// 		Val: 4,
	// 		Left: &TreeNode{
	// 			Val: 1,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 2,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 5,
	// 	},
	// }
	// B = &TreeNode{
	// 	Val: 4,
	// 	Left: &TreeNode{
	// 		Val: 1,
	// 	},
	// }
	t.Log(isSubStructure(A, B))
}
