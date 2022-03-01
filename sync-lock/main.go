package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	num int64
	rw  sync.RWMutex
	wg  sync.WaitGroup
)

func write() {
	defer wg.Done()
	rw.Lock()
	num = num + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	fmt.Printf("write data: %v\n", num)
	rw.Unlock()
	fmt.Println("write over")
}

func read() {
	defer wg.Done()
	rw.RLock()                   // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	fmt.Printf("Read data: %v\n", num)
	rw.RUnlock() // 解读锁
	fmt.Println("Read over")

}

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println("num:", num)

}
