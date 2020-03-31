package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	repeat := func(done chan interface{},value int) <-chan int{
		valueStream := make(chan int)
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- value:
					time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
				}
			}
		}()

		return valueStream
	}

	take := func(done chan interface{} , valueStream <-chan int , num int) <-chan int{

		takeStream := make(chan int)
		go func() {
			defer close(takeStream)
			for i:=0;i<num;i++{
				select {
				case <-done:
					return
				case takeStream <- i+<-valueStream:

				}
			}
		}()
		return takeStream
	}
	

	done := make(chan interface{})
	defer close(done)

	for v := range merge(take(done,repeat(done , 1),10),take(done,repeat(done , 10),10)){
		fmt.Println(v)
	}

}

func merge(a <-chan int, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for a != nil|| b != nil{
			select {
			case v , ok := <- a:
				if !ok {
					a = nil
					fmt.Println("a channel is done")
					continue
				}
				c <- v
			case v , ok:= <- b:
				if !ok {
					b = nil
					fmt.Println("b channel is done")
					continue
				}
				c <- v
			}
		}
	}()

	return c
}
