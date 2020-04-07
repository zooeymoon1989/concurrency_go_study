package main

import "fmt"

func main() {

	doWork := func(done chan interface{}, strings chan string) chan interface{} {

		completed := make(chan interface{})
		go func() {
			defer close(completed)
			defer fmt.Println("go routine is done")
			for {
				select {
				case <-done:
					return
				case completed <-strings:
				}
			}
		}()
		return completed
	}

	done := make(chan interface{})
	complete := doWork(done, nil)
	go func() {
		defer close(done)

	}()
	<-complete
	fmt.Println("done")

}
