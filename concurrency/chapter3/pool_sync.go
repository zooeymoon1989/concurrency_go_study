package main

import (
	"fmt"
	"sync"
)

func main() {
	//myPool := &sync.Pool{New: func() interface{} {
	//	fmt.Println("Creating new instance")
	//	return struct {}{}
	//}}
	//
	//myPool.Get()
	//instance := myPool.Get()
	//myPool.Put(instance)
	//myPool.Get()

	//////////////////////////////////////

	var numCalcsCreated int
	calcPool := sync.Pool{New: func() interface{} {
		numCalcsCreated++
		mem := make([]byte, 1024)
		return &mem
	}}

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorks = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorks)
	for i := numWorks; i > 0; i-- {

		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			calcPool.Put(mem)
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)

	///////以上结果输出12，因为有gc，所以实例化的对象会被自动清理

}
