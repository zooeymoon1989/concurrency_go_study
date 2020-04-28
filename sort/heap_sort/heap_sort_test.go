package heap_sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {

	array := []int{2,4,8,9,7,10,9,15,20,13}
	result := HeapSort(array)
	fmt.Println(result)
}
