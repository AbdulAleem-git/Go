package main

import (
	"context"
	"fmt"
	"time"
)

func processRequest(ctx context.Context) {
	// Perform some time-consuming task
	for i := 1; i <= 15; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Processing step", i)
		case <-ctx.Done():
			fmt.Println("Operation canceled")
			return
		}
	}
	fmt.Println("Processing completed")
}

func main() {
	// Create a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Execute the task in a separate goroutine
	go processRequest(ctx)

	// Wait for the context timeout or cancellation signal
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Main goroutine completed")
	case <-ctx.Done():
		fmt.Println("Main goroutine canceled due to timeout")
	}

	time.Sleep(*time.Second)
}
