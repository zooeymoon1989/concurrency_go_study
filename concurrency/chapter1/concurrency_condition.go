package chapter1

import (
	"fmt"
	"sync"
)

func ConcurrencyCondition() {

	var memoryAccess sync.Mutex

	var i int

	go func() {
		fmt.Println("goroutine is running")
		memoryAccess.Lock()
		i++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if i == 0 {
		fmt.Println("the value is 0")
	} else {
		fmt.Printf("the value is %v.\n", i)
	}
	fmt.Println("no goroutine is running")
	memoryAccess.Unlock()
}
