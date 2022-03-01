package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter interface {
	Inc()
	Load() int64
}

type CommonCounter struct {
	counter int64
}

func (c CommonCounter) Inc() {
	c.counter++
}
func (c CommonCounter) Load() int64 {
	return c.counter
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			fmt.Println(c.Load())
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
func main() {
	c1 := CommonCounter{}
	test(c1)
	fmt.Println(c1)


}
