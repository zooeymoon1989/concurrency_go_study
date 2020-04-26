package select_sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	array := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := SelectSort(array)
	fmt.Println(result)
}
