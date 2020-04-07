package algorithm

import "math"

/*

	Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

	Find all the elements that appear twice in this array.

	Could you do it without extra space and in O(n) runtime?

 */


func FindAllDuplicatesInAnArray(input ...int) (output []int) {
	for _, v := range input {
		tmpValue := int(math.Abs(float64(v))) - 1
		if input[tmpValue] < 0 {
			output = append(output, tmpValue+1)
		}
		input[tmpValue] = -input[tmpValue]
	}

	return
}
