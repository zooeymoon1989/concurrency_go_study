package quick_sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := []int{0,3,2,5,12,-1,-23,-4,-5,9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := QuickSort(array)
	fmt.Println(result)
}
