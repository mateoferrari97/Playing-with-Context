package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		/*
			1. Listen stdin and cancel ctx if so.
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()

			OR...

			2. After 2 seconds, cancel ctx.
			time.Sleep(2 * time.Second)
		*/


		cancel()
	}()

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
