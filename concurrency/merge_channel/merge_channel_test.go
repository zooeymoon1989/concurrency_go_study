package main

import (
	"fmt"
	"testing"
)

var funcs = []struct {
	name string
	f    func(...<-chan int) <-chan int
}{
	{
		"goroutines",
		MergeChan,
	}, {
		"reflection",
		MergeReflect,
	},
}

func BenchmarkMergeChan(b *testing.B) {

	for _ , v := range funcs {
		for n := 1; n <= 1024; n *= 2 {

			b.Run(fmt.Sprintf("%s/%d", v.name , n), func(b *testing.B) {
				chans := make([]<-chan int, n)
				for i := 0; i < b.N; i++ {

					for i := range chans {
						chans[i] = AsMergeChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
					}

					c := v.f(chans...)

					for range c {

					}

				}

			})
		}
	}
}
