package context

import (
	"context"
	"fmt"
	"time"
)

func Context() {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	mySleepAndTalk(ctx, time.Second*5, "hello")

}

func mySleepAndTalk(ctx context.Context, d time.Duration, s string) {
	select {
	case <-time.After(d):
		fmt.Println(s)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func sleepAndTalk(ctx context.Context, duration time.Duration, s string) {

	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())

	case <-time.After(duration):
		fmt.Println(s)
	}
}
