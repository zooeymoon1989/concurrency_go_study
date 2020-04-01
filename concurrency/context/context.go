package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second) // 需要传递父context

	defer cancel()

	mySleepAndTalk(ctx, time.Second*5, "hello")

}

func mySleepAndTalk(ctx context.Context, d time.Duration, s string) {
	time.Sleep(d)
	fmt.Println(s)
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
