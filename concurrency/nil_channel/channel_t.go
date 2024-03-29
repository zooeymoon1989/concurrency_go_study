package nil_channel

import "fmt"

func callerA(c chan<- string) {
	c <- "hello world"
	close(c)
}

func callerB(c chan<- string) {
	c <- "hola "
	close(c)
}

func ChannelT() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 {
		select {
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}else{
				a = nil
				fmt.Println("a")
			}

		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from B\n", msg)
			}else{
				b = nil
				fmt.Println("b")
			}

		}
	}
}
