package main

import "fmt"

func main() {

	orDone := func(done, c <-chan interface{}) <-chan interface{} {

		valStream := make(chan interface{})

		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if !ok {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()

		return valStream
	}

	tee := func(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {

		out1, out2 := make(chan interface{}), make(chan interface{})

		go func() {
			defer close(out1)
			defer close(out2)

			for val := range orDone(done, in) {
				var out1, out2 = out1, out2  //没有明白什么意思，为什么要隐藏外部的out1和out2变量
				for i := 0; i < 2; i++ {
					select {
					case <-done:
						return
					case out1 <- val:
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}

		}()
		return out1, out2
	}

	repeat := func(done <-chan interface{}, values ...int) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, i := range values {
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

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:

				}
			}
		}()

		return takeStream
	}

	done := make(chan interface{})
	defer close(done)
	out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4))
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}

}
