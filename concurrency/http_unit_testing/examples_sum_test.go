package sum_test

import (
	"fmt"
	sum "github.com/my/repo/concurrency/http_unit_testing"
)

func ExampleInts() {
	s := sum.Ints(1,2,3,4,5)
	fmt.Printf("sum of one to five is %v",s)
	// Output:
	// sum of one to five is 15
}