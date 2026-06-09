package main

import (
	"context"
	"fmt"
	"time"
)

func Worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped")
			return
		default:
			fmt.Println("working...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go Worker(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(time.Second)
}
