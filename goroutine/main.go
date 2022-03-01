package main

import (
	"fmt"
	"time"
)

func test(ch chan int) {
	go func() {
		defer func() {
			fmt.Printf("defer...")
		}()

		for {
			select { // 阻塞等待任务、结束信号到来
			case task, ok := <-ch: // 从 channel 中消费任务
				if !ok { // 如果 channel 被关闭, 结束 worker 运行
					return
				}
				// 执行任务
				fmt.Println("开始..", task)

			}
		}
	}()
}
func main() {
	ch := make(chan int, 20)
	test(ch)

	time.Sleep(1 * time.Second)

	for i := 0; i < 20; i++ {
		ch <- i
	}
	time.Sleep(4 * time.Second)
}
