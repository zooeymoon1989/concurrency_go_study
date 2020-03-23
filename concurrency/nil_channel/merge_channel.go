package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	a := asMergeChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := asMergeChan(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	c := asMergeChan(20, 21, 22, 23, 24, 25, 26, 27, 28, 29)
	d := asMergeChan(30, 31, 32, 33, 34, 35, 36, 37, 38, 39)

	for v := range mergeChan(a, b, c, d) {
		fmt.Println(v)
	}

}

func mergeChan(chans ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chans))

		for _, c := range chans {
			go func(<-chan int) {
				for v := range c {
					out <- v
				}
				wg.Done()
			}(c)
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func asMergeChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()

	return c
}
