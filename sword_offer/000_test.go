package sword_offer

import "testing"

/**
* @Author XuPEngHao
* @Date 2023/6/5 11:04
 */
func Test_003(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	nums = []int{3, 4, 2, 1, 1, 0}
	t.Log(findRepeatNumberV3(nums))
}
