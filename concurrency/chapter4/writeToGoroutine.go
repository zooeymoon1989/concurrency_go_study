package main

//func main() {
//	newRandStream := func(done <-chan interface{}) <-chan int {
//
//		randStream := make(chan int)
//		go func() {
//			defer fmt.Println("newRandStream closure exited!")
//			defer close(randStream)
//			for {
//				select {
//				case randStream <- rand.Int():
//				case <-done:
//					return
//				}
//			}
//		}()
//		return randStream
//	}
//
//	done := make(chan interface{})
//	randStream := newRandStream(done)
//	fmt.Println("3 random ints:")
//
//	for i := 1; i <= 3; i++ {
//		fmt.Printf("%d:%d\n", i, <-randStream)
//	}
//	close(done) //这个close是用来发射信号到newRandStream中来函数退出，从而防止内存泄漏
//
//	time.Sleep(1 * time.Second)
//}
