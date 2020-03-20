package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func main() {

	db, _ := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())

	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
