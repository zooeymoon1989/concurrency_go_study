package server

import (
	"context"
	"fmt"
	"github.com/my/repo/concurrency/context/log"
	"net/http"
	"time"
)

func Server() {
	http.HandleFunc("/", log.Decorate(serverHandler))
	panic(http.ListenAndServe("127.0.0.1:8080", nil))
}

func serverHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	ctx = context.WithValue(ctx , int(42) , int64(100))
	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler stopped")
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello world")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
