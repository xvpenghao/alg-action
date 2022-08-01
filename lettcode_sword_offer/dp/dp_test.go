package dp

import "testing"

// @Author XuPEngHao
// @DATE 2022/7/5 21:21

func Test_maxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	t.Log(maxSubArray(nums))
}

func Test_fib(t *testing.T) {
	t.Log(fib(5))
}

func Test_numWays(t *testing.T) {
	t.Log(numWays(7))
}
