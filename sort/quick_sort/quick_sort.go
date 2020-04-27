package quick_sort

import "math/rand"

// 快速排序
//
// 快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。
//
// 1.从数列中挑出一个元素，称为 “基准”（pivot）；
// 2.重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
// 3.递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。

func QuickSort(array []int) []int {

	return quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, left, right int) []int {
	if left < right {
		pivot := partition(array , left , right)
		quickSort(array,left , pivot-1)
		quickSort(array , pivot+1 , right)
	}

	return array
}

func partition(array []int, left, right int) int {

	pivot := getPivot(right - left + 1 , left)
	swap(array , left , pivot)

	index := left + 1


	for i := index ; i <= right ; i ++ {

		if array[i] < array[left] {
			swap(array , i , index)
			index++
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
