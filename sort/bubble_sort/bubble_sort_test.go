package bubble_sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	//array := []int{ 9,8,7,6,5,4,3,2,1}
	array := []int{ 34,-2,54,54,-23,4,12,6,23}
	result := BubbleSort(array)
	fmt.Println(result)
}
