package pointer

import (
	"fmt"
	"testing"
)

func TestFindSumInArray(t *testing.T) {

	sum := 15
	array := []int{1, 2, 4, 7, 11, 15}

	result := FindSumInArray(array, sum)
	if result == nil {
		t.Errorf("could not find two numbers sum equals %d", sum)
	}

}

func TestFindSequenceArray(t *testing.T) {
	sum := 15
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := FindSequenceArray(array, sum)
	if result == nil {
		t.Errorf("could not find two numbers sum equals %d", sum)
	}

}

func TestFindSequenceArrayFromNumber(t *testing.T) {
	sum := 66

	result := FindSequenceArrayFromNumber(sum)

	fmt.Println(result)

}
