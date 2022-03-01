package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func work(ctx context.Context) {
LOOP:
	for {
		time.Sleep(time.Microsecond * 5)
		select {
		case <-ctx.Done():
			fmt.Println("work time out") //假设数据库链接需要5毫秒
			break LOOP
		default:
			fmt.Println("work..") //假设数据库链接需要5毫秒
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}
func main() {
	//now := time.Now().Add(time.Second * 10)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	wg.Add(1)
	go work(ctx)

	time.Sleep(time.Second * 5)

	cancel()
	wg.Wait()
	fmt.Println("over")

}
