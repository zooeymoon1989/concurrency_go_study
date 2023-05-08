package chapter3

import (
	"fmt"
	"sync"
	"time"
)

func SelectChannel() {

	//start := time.Now()
	//c := make(chan interface{})
	//
	//go func() {
	//	time.Sleep(5*time.Second)
	//	close(c)
	//}()
	//fmt.Println("Blocking on read...")
	//select {
	//case <-c:
	//	fmt.Printf("Unblocked %v later.\n", time.Since(start))
	//}

	//
	//c := make(chan int)
	//go func() {
	//	c <- 1
	//	defer close(c)
	//}()
	//
	//select {
	//case <-c :
	//	fmt.Println(c)
	//case <-time.After(1*time.Second):
	//	fmt.Println("Timed out.")
	//
	//}

	//var c = make(chan string)
	//var d = make(chan string)
	//
	//a := func(c chan<- string) {
	//	c <- "Hello World!"
	//}
	//b := func(c chan<- string) {
	//	c <- "Hola Mundo!"
	//}
	//
	//go a(c)
	//go b(d)
	//for i := 0; i < 5; i++ {
	//	select {
	//	case msg := <-c:
	//		fmt.Printf("%s  from A\n", msg)
	//	case msg := <-d:
	//		fmt.Printf("%s from B\n", msg)
	//	case <-time.After(1 * time.Second):
	//		fmt.Println("times out")
	//	}
	//}

	//c1 := make(chan interface{});close(c1)
	//c2 := make(chan interface{});close(c2)
	//
	//var c1count , c2count int
	//
	//for i:=0;i<1000;i++{
	//	select {
	//	case <-c1:
	//		c1count++
	//	case <-c2:
	//		c2count++
	//	}
	//}
	//fmt.Printf("c1Count: %d\nc2Count: %d\n", c1count, c2count)

	//var c <-chan int
	//select {
	//case <-c:
	//case <-time.After(1 * time.Second):
	//	fmt.Println("timed out.")
	//}

	//var some int32
	begin := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//atomic.AddInt32(&some, 1)
			begin <- i
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")

	for i := 0; i < 5; i++ {

		select {
		case msg, ok := <-begin:
			if ok {
				fmt.Println(msg)
			}
		case <-time.After(1 * time.Second):
			fmt.Println("times out")
		default:
			fmt.Println("default")
		}

	}

	wg.Wait()

	//	done := make(chan interface{})
	//
	//	go func() {
	//		time.Sleep(5 * time.Second)
	//		close(done)
	//	}()
	//	workCounter := 0
	//loop:
	//	for {
	//		select {
	//		case <-done:
	//			break loop
	//		default:
	//
	//		}
	//
	//		workCounter++
	//		time.Sleep(1 * time.Second)
	//	}
	//
	//	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)

}
