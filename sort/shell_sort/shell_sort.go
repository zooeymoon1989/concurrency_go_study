package shell_sort

func ShellSort(array []int) []int {

	if len(array) <= 1 {
		return array
	}

	for j := len(array) / 2; j > 0; j = j / 2 {
		for i := j; i < len(array); i++ {
			k := i - j
			current := array[k]

			for k >= 0 && array[k] > array[k+j] {
				array[k+j] = array[k]
				k = k - j
			}

			array[k+j] = current

		}

	}

	return array

}
