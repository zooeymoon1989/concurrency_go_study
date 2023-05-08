package reorder_array

func ReorderArray(array []int) []int {

	if len(array) <= 1 {
		return array
	}

	last := len(array)
	p1 := 0

	for p1 < last {
		if array[p1]%2 == 0 {

			current := array[p1]
			for j := p1; j < len(array)-1; j++ {
				array[j] = array[j+1]
			}
			array[len(array)-1] = current
			last--
		}else {
			p1++
		}
	}
	return array
}
