package main

import (
	"context"
	"log"
	"time"
)

/*
	Instead of instantiate context.WithCancel() and set a Timer with time.AfterFunc() to call cancel() after t seconds.
	We can use context.WithTimeout() to do the same.
*/

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	talkAfter(ctx, 6*time.Second, "Holis")
}

func talkAfter(ctx context.Context, after time.Duration, msg string) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
	case <-time.After(after):
		log.Println(msg)
	}
}
