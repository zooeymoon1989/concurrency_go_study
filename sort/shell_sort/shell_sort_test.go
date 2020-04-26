package shell_sort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	array := []int{9,8,7,6,5,4,3,2,1}
	result := ShellSort(array)
	fmt.Println(result)
}
