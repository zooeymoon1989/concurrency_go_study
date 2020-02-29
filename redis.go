package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func main() {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{":26379"},
	})

	rdb.Ping()

	//rdb.Set("foo","bar" ,0)

	fmt.Println(rdb.Get("foo").Val())

}
