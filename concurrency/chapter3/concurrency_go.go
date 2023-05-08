package chapter3

import (
	"fmt"
	"sync"
)

func ConcurrencyGo() {
	var wg sync.WaitGroup
	for _, salutation := range []int{1, 2, 3} {
		wg.Add(1)
		go func(salutation int) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
