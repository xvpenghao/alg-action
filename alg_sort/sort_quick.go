package alg_sort

/**
* @Author XuPEngHao
* @Date 2023/6/15 07:27
 */

func quickSortV2(numList []int, left, right int) {
	if left < right {
		i, j := left, right
		tmpX := numList[i]
		for i < j {
			for i < j && numList[j] <= tmpX {
				j--
			}
			if i < j {
				numList[i] = numList[j]
				i++
			}
			for i < j && numList[i] > tmpX {
				i++
			}
			if i < j {
				numList[j] = numList[i]
				j--
			}
		}
		numList[j] = tmpX
		quickSortV2(numList, left, i-1)
		quickSortV2(numList, i+1, right)
	}
}

func quick(numList []int, left, right int) {
	if left < right {
		i := quickSort(numList, left, right)
		quickSort(numList, left, i-1)
		quickSort(numList, i+1, right)
	}
}

func quickSort(numList []int, left, right int) int {
	tmpX := numList[left]
	for left < right {
		for left < right && numList[right] >= tmpX {
			right--
		}
		if left < right {
			numList[left] = numList[right]
			left++
		}

		for left < right && numList[left] < tmpX {
			left++
		}
		if left < right {
			numList[right] = numList[left]
			right--
		}
	}

	numList[left] = tmpX
	return left
}
