package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	http.HandleFunc("/", Decorate(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		Log(ctx, "handle started")
		Log(ctx, "handle ended")

		select {
		case <-time.After(5 * time.Second):
			_, _ = fmt.Fprintln(w, "Holis")
		case <-ctx.Done():
			Log(ctx, ctx.Err().Error())
			http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
		}
	}))

	return http.ListenAndServe(":8080", nil)
}

type key int

const ctxKey = key(1)

func Decorate(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := rand.Int63()
		ctx = context.WithValue(ctx, ctxKey, requestID)

		h(w, r.WithContext(ctx))
	}
}

func Log(ctx context.Context, msg string) {
	requestID := ctx.Value(ctxKey).(int64)
	log.Printf("[%d]: %s\n", requestID, msg)
}
