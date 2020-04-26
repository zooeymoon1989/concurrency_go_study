package shell_sort

// 希尔排序
// 1959年Shell发明，第一个突破O(n2)的排序算法，是简单插入排序的改进版。
// 它与插入排序的不同之处在于，它会优先比较距离较远的元素。希尔排序又叫缩小增量排序。
//
// 实现步骤：
// 1.选择一个增量序列t1，t2，…，tk，其中ti>tj，tk=1；
// 2.按增量序列个数k，对序列进行k 趟排序；
// 3.每趟排序，根据对应的增量ti，将待排序列分割成若干长度为m 的子序列，分别对各子表进行直接插入排序。仅增量因子为1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
//
func ShellSort(array []int) []int {

	if len(array) <= 1 {
		return array
	}

	for j := len(array) / 2; j > 0; j = j / 2 {

		for i := j;i < len(array) ; i++ {

			k := i
			current := array[k]

			for k - j >= 0 && array[k-j] > current {

				array[k] = array[k-j]
				k = k - j
			}

			array[k] = current

		}

	}

	return array

}
