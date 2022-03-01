package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//context 处理单个请求中多个goroutine数据与请求域数据， 取消信号，截止时间等相关操作

//context 接口， 包含4个需要实现的方法

//Deadline， Done ， Err， Value
//Deadline： 返回当前goroutine 被取消的时间
//Done： 返回的是个channel， 这个channel当工作完成或者被取消时候关闭，多次调用Done会返回用一个channel
//Err：返回Context结束的原因
//Value ：该方法用于传递跨API和进程跟请求域的数据

var wg sync.WaitGroup

func work(ctx context.Context) {
	go work2(ctx)
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println("work...")
		}
	}
	wg.Done()
}

func work2(ctx context.Context) {
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println("work2...")
		}
	}
}
func main() {

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go work(ctx)
	time.Sleep(time.Second)
	cancel()

	wg.Wait()
	fmt.Println("exit...")
}
