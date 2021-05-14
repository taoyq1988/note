package sort

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
