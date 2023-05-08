package chapter4

import "fmt"

//import "fmt"
//
//func main() {
//	//这里使用了一个闭包，把channel放入函数中，这样可以防止其他goroutine写入
//	chanOwner := func() <-chan int {
//		results := make(chan int, 5)
//		go func() {
//			defer close(results)
//			for i := 0; i < 5; i++ {
//				results <- i
//			}
//		}()
//
//		return results
//	}
//
//	// 这里我们的参数设置成一个只读
//	consumer := func(results <-chan int) {
//		for result := range results {
//			fmt.Printf("Received: %d\n", result)
//		}
//		fmt.Println("Done receiving!")
//	}
//
//	//在闭包内的channel发送给消费者
//	consumer(chanOwner())
//
//}

func Closure() {
	chanOwner := func() <-chan int{
		r := make(chan int,5)
		go func() {
			defer close(r)
			for i:=0;i<=5;i++{
				r <- i
			}
		}()

		return r
	}

	comsumer := func(r <-chan int) {

		for v := range r {
			fmt.Println(v)
		}
		fmt.Println("finished")
	}

	comsumer(chanOwner())
}


