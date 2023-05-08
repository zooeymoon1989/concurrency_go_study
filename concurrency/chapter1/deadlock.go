package chapter1

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int8
}

func DeadLock() {

	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	printSum(&a, &b)
	printSum(&b, &a)
	wg.Wait()

}
