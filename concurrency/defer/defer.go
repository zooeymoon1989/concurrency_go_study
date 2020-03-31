package main

import "fmt"

func main() {
	fmt.Printf("with named param, x: %d\n", namedParam())
	fmt.Printf("without named param, x: %d\n", notNamedParam())
}
func namedParam() (x int) {
	x = 1
	defer func() { x = 2 }()
	return x
}

func notNamedParam() int {
	x := 1
	defer func() { x = 2 }()
	return x
}
