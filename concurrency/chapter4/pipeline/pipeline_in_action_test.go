package pipeline_in_action

import "testing"

func BenchmarkGeneric(b *testing.B) {

	take := func(done <-chan interface{}, valueStream interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i <= num; i++ {
				select {
				case <-done:
					return
				case takeStream <- valueStream:

				}
			}
		}()

		return takeStream
	}

	done := make(chan interface{})
	defer close(done)
	b.ResetTimer()
	for range ToString(done, take(done, "a", b.N)) {

	}
}

func BenchmarkTyped(b *testing.B) {

	repeat := func(done <-chan interface{}, values ...string) <-chan string {
		valueStream := make(chan string)
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()

		return valueStream
	}

	take := func(
		done <-chan interface{}, valueStream <-chan string, num int,
	) <-chan string {
		takeStream := make(chan string)
		go func() {
			defer close(takeStream)
			for i := num; i > 0 || i == -1; {
				if i != -1 {
					i--
				}
				select { case <-done:
					return
				case takeStream <- <-valueStream: }
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer  close(done)

	b.ResetTimer()
	for range take(done , repeat(done , "a") , b.N){

	}

}

func ToString(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case stringStream <- v.(string):

			}
		}

	}()
	return stringStream
}
