package insert_sort

import (
	"fmt"
	"testing"
)

func TestInsertSorted(t *testing.T) {
	array := []int{32,-1,1,2,-45,-123}
	result := InsertSorted(array)
	fmt.Println(result)
}
