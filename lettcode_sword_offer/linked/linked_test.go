package linked

import "testing"

// @Author XuPEngHao
// @DATE 2022/5/3 10:54

func Test_getIntersectionNode(t *testing.T) {
	l1 := &ListNode{
		Val: 4, Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 5, Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}
	node := getIntersectionNode(l1, l2)
	t.Log(node)
}

func Test_mergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	node := mergeTwoLists(l1, l2)
	t.Log(node)
}

func Test_reverseList(t *testing.T) {
	node := reverseList(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	})
	t.Log(node)
}

func Test_getKthFromEnd(t *testing.T) {
	node := getKthFromEnd(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}, 6)
	t.Log(node)
}

func Test_deleteNode(t *testing.T) {
	node := deleteNodeV2(&ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}, 5)
	t.Log(node)
}
