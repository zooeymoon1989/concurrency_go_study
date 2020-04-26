package insert_sort

// 插入排序
// 插入排序（Insertion-Sort）的算法描述是一种简单直观的排序算法。
// 它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
//
// 实现步骤：
// 1。从第一个元素开始，该元素可以认为已经被排序；
// 2。取出下一个元素，在已经排序的元素序列中从后向前扫描；
// 3。如果该元素（已排序）大于新元素，将该元素移到下一位置；
// 4。重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
// 5。将新元素插入到该位置后；
// 6。重复步骤2~5。
func InsertSorted(array []int) []int {

	if len(array) == 0 {
		return nil
	}

	if len(array) == 1 {
		return array
	}

	for j := 1; j < len(array); j++ {
		i := j - 1
		for i >= 0 && array[i] > array[i+1] {
			tmp := array[i]
			array[i] = array[i+1]
			array[i+1] = tmp
			i--
		}

	}

	return array

}
