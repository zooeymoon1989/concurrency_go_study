package heap_sort

// 堆排序
// 堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。堆排序可以说是一种利用堆的概念来排序的选择排序。分为两种方法：
//
// 大顶堆：每个节点的值都大于或等于其子节点的值，在堆排序算法中用于升序排列；
// 小顶堆：每个节点的值都小于或等于其子节点的值，在堆排序算法中用于降序排列；
// 堆排序的平均时间复杂度为 Ο(nlogn)。
//
// 实现步骤：
// 1.创建一个堆 H[0……n-1]；
// 2.把堆首（最大值）和堆尾互换；
// 3.把堆的尺寸缩小 1，并调用 shift_down(0)，目的是把新的数组顶端数据调整到相应位置；
// 4.重复步骤 2，直到堆的尺寸为 1。
//
func HeapSort(array []int) []int {
	length := len(array)
	buildHeap(array, length)

	for i := length - 1; i >= 0; i-- {
		swap(array, 0, i)
		length--
		heapify(array, 0, length)
	}

	return array
}

func buildHeap(array []int, length int) {
	for i := length / 2; i >= 0; i-- {
		heapify(array, i, length)
	}
}

func heapify(array []int, index, length int) {
	left := index*2 + 1
	right := index*2 + 2

	current := index

	if left < length && array[left] > array[current] {
		current = left
	}

	if right < length && array[right] > array[current] {
		current = right
	}

	if current != index {
		swap(array, current, index)
		heapify(array, current, length)
	}

}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
