package quick_sort

import "math/rand"

func QuickSort(array []int) []int {

	return quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, left, right int) []int {
	if left < right {
		index := partition(array, left, right)
		quickSort(array, left, index-1)
		quickSort(array, index+1, right)
	}
	return array
}

func partition(array []int, left, right int) int {

	pivot := getPivot(right-left+1, left)
	index := left + 1
	swap(array, left, pivot)

	for i := index; i <= right; i++ {
		if array[i] < array[left] {
			swap(array, i, index)
			index ++
		}
	}

	swap(array , left , index - 1)
	return index - 1
}

// 生成一个基准(pivot)
func getPivot(length, left int) int {
	return rand.Intn(length) + left
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
