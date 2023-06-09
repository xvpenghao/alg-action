package hot_100

import "testing"

/**
* @Author XuPEngHao
* @Date 2023/6/9 10:57
 */

func Test_001(t *testing.T) {
	numsList := []int{2, 7, 11, 15}
	target := 9

	numsList = []int{3, 2, 4}
	target = 6

	t.Log(twoSum2(numsList, target))
}
