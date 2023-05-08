package pipeline_in_action

import (
	"fmt"
)

func PipeLineInAction() {
	repeat := func(done <-chan interface{},values... int) <-chan interface{}{
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _ , i := range values {
					select {
					case <-done:
						return
					case valueStream <- i:
					}
				}
			}
		}()

		return valueStream
	}

	//repeatFn := func(done chan interface{}, fn func() interface{}) chan interface{} {
	//	valueStream := make(chan interface{})
	//	go func() {
	//		defer close(valueStream)
	//		for {
	//			select {
	//			case <-done:
	//				return
	//			case valueStream <- fn():
	//			}
	//		}
	//
	//	}()
	//
	//	return valueStream
	//}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i <= num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:

				}
			}
		}()

		return takeStream
	}



	//randNum := func() interface{} {
	//	return rand.Int()
	//}

	done := make(chan interface{})
	defer close(done)

	for num := range take(done , repeat(done , 1,2,3,4,5,6,7,8,9,10) , 9) {
		fmt.Println(num)
	}

	//for num := range take(done, repeatFn(done, randNum), 10) {
	//	fmt.Println(num)
	//}

}
