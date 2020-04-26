package merge_sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	array := []int{9,8,7,6,5,4,3,2,1}
	result := MergeSort(array)
	fmt.Println(result)
}
