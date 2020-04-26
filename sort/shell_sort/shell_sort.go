package shell_sort

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
