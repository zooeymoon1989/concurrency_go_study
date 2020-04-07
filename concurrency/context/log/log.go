package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type key int

const RequestIDKey = key(42)

func Println(ctx context.Context , msg string) {

	id, ok := ctx.Value(RequestIDKey).(int64)
	if !ok {
		log.Println("could not find request ID in context")
	}

	log.Printf("[%d] %s" , id , msg)
}

func Decorate(f http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx , RequestIDKey , id)
		f(w , r.WithContext(ctx))
	}
}
