package merge_sort

// 归并排序
// 归并排序是建立在归并操作上的一种有效的排序算法。
// 该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
// 将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。
//
// 实现步骤：
// 1.把长度为n的输入序列分成两个长度为n/2的子序列；
// 2.对这两个子序列分别采用归并排序；
// 3.将两个排序好的子序列合并成一个最终的排序序列。
func MergeSort(array []int) []int {

	if len(array) < 2 {
		return array
	}

	left := MergeSort(array[0 : len(array)/2])
	right := MergeSort(array[len(array)/2:])
	return merge(left, right)
}

func merge(left, right []int) []int {

	arr := make([]int, len(left)+len(right))

	l, r, i := 0, 0, 0

	for {

		if left[l] > right[r] {
			arr[i] = right[r]
			r++
			i++

			if r == len(right) {
				copy(arr[i:], left[l:])
				break
			}
		} else {
			arr[i] = left[l]
			l++
			i++

			if l == len(left) {
				copy(arr[i:], right[r:])
				break
			}
		}

	}

	return arr
}
