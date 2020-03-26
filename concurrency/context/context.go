package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	
	ctx := context.Background()

	ctx , cancel := context.WithCancel(ctx) // 需要传递父context

	// 协程
	// 扫描任何输入
	// 结束context
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	sleepAndTalk(ctx , time.Second*5,"hello")
	
}

func sleepAndTalk(ctx context.Context, duration time.Duration, s string) {
	c, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	select {
	case <-c.Done():
		fmt.Println("handle", ctx.Err())

	case <-time.After(duration):
		fmt.Println(s)
	}
}
