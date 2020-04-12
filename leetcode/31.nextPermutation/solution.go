package _1_nextPermutation

func nextPermutation(nums []int) {
	i := len(nums)-1
	for ; i>=1 ; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}
	i--
	j := i+1
	if i == -1 {
		QuickSort(nums)
		return
	}
	for ; j<len(nums); j++ {
		if nums[j] <= nums[i] {
			break
		}
	}
	nums[i], nums[j-1] = nums[j-1], nums[i]
	QuickSort(nums[i+1:])
}

// quick sort
func QuickSort(array []int) {
	if len(array) == 0 {
		return
	}
	left := 0
	right := len(array) - 1
	obj := array[left]
	for left < right {
		for left < right && array[right] >= obj {
			right--
		}
		array[left] = array[right]

		for left < right && array[left] <= obj {
			left++
		}
		array[right] = array[left]
	}
	array[left] = obj
	QuickSort(array[:left])
	QuickSort(array[right+1:])
}
