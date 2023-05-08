package reorder_array

import (
	"fmt"
	"testing"
)

func TestReorderArray(t *testing.T) {
	//array := []int{1,2,3,4,5,6,7}
	array := []int{2,4,6,1,3,5,7}
	result := ReorderArray(array)
	fmt.Println(result)
}
