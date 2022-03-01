package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	now := time.Now().Add(10 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), now)

	defer cancel()

	select {
	case <-time.After(time.Second):
		fmt.Println("over")
	case <-ctx.Done():
		fmt.Println("context err:", ctx.Err())
	}
}
