package chapter3

import (
	"bytes"
	"fmt"
	"os"
)

func Channel() {

	//begin := make(chan interface{})
	//var wg sync.WaitGroup
	//for i:=0;i<5;i++{
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		<-begin
	//		fmt.Printf("%v has begun\n", i)
	//	}(i)
	//}
	//
	//fmt.Println("Unblocking goroutines...")
	//close(begin)
	//wg.Wait()

	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)

	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}

}
