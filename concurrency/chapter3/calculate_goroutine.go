package chapter3

import (
	"fmt"
	"runtime"
	"sync"
)

func CalculateGoroutines() {

	//在使用一次gc后返回系统总共的内存
	memConsume := func() uint64 {
		runtime.GC() //执行一次gc
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c } //不会退出的goroutine

	const numGoroutines = 100000
	wg.Add(numGoroutines)
	before := memConsume() // 创建goroutine之前的总消耗内存

	for i := numGoroutines; i > 0; i-- {
		go noop()
	}

	wg.Wait()
	after := memConsume() // 创建goroutine之后的总消耗内存
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)

}
