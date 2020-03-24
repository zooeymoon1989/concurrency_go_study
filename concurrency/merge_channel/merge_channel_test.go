package main

import (
	"fmt"
	"testing"
)

func BenchmarkMergeChan(b *testing.B)  {

	for i := 1 ; i <= 1024 ; i *= 2 {
		b.Run(fmt.Sprintf("%d",i), func(b *testing.B) {
			c := MergeReflect(AsMergeChan(1,2,3), AsMergeChan(4,5,6), AsMergeChan(7,8,9))
			seen:=make(map[int]bool)
			for v := range c {
				if seen[v] {
					b.Errorf("saw %d at least twice",v)
				}
				seen[v] = true
			}

			for i := 1;i<=9;i++{
				if !seen[i] {
					b.Errorf("didn't see %d",i)
				}
			}
		})
	}
}
