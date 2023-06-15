package alg_sort

// @Author XuPEngHao
// @DATE 2022/3/9 09:07

// 什么叫 冒泡排序,
// 每一轮的排序，都会确定好这个值的位置
func bubble(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
