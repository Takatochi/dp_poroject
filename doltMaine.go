package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// create a new context with a timeout of 1 second
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(ctx)

	//create a new context with deadline of 1 second
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel = context.WithDeadline(context.Background(), deadline)
	defer cancel()
	fmt.Println(ctx)
}
