package fan_in_fan_out

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func FanInFanOut() {


	fanIn := func(done chan interface{} , channels ...<-chan interface{}) <-chan interface{}{

		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()

			for i:= range c{
				select {
				case <-done:
					return
				case multiplexedStream <- i:

				}
			}

		}

		wg.Add(len(channels))

		for _ , c := range  channels{
			go multiplex(c)
		}

		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()

		return multiplexedStream
	}

	repeatFn := func(done chan interface{}, fn func() interface{}) chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}

		}()

		return valueStream
	}

	randInt := func() interface{} {
		return rand.Intn(50000000)
	}

	toInt := func(done <-chan interface{}, valueStream chan interface{}) chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()

		return intStream

	}

	findPrime := func(done <-chan interface{}, intStream <-chan int) <-chan interface{} {

		primeStream := make(chan interface{})
		go func() {
			defer close(primeStream)

			for integer := range intStream {

				integer -= 1
				prime := true

				for divisor := integer - 1; divisor > 1; divisor-- {

					if integer%divisor == 0 {
						prime = false
						break
					}

				}

				if prime {
					select {
					case <-done:
						return
					case primeStream <- integer:
					}
				}
			}
		}()

		return primeStream
	}


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

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done , repeatFn(done,randInt))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{} , numFinders)


	fmt.Println("Primes:")

	for i:=0;i<numFinders;i++{
		finders[i] = findPrime(done,randIntStream)
	}

	for prime := range take(done , fanIn(done,finders...) , 20) {
		fmt.Printf("\t%d\n", prime)
	}


	fmt.Printf("Search took: %v", time.Since(start))

}
