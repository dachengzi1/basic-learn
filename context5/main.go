package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

var wg sync.WaitGroup

func work(ctx context.Context) {
	key := TraceCode("TRACE_CODE")

	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		time.Sleep(time.Microsecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println("work ....", traceCode)
		}
	}
	fmt.Println("go done ....")
	wg.Done()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)

	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "222")
	wg.Add(1)
	go work(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
