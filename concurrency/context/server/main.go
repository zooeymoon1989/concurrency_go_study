package server

import (
	"context"
	"fmt"
	"github.com/my/repo/concurrency/context/log_server"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", log_server.Decorate(handler))
	panic(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx , 42 , int64(250))

	log_server.Println(ctx, "handler started")
	defer log_server.Println(ctx, "handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		log_server.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
