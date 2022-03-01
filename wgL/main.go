package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var rw sync.RWMutex
var x int64
func write()  {
	rw.Lock()
	x = x+ 1
	time.Sleep(time.Microsecond * 10)
	rw.Unlock()
	wg.Done()
}

func read()  {
	rw.RLock()
	time.Sleep(time.Microsecond)
	rw.RUnlock()
	wg.Done()
}
func main()  {
	start:=time.Now()

	for i:=0; i<10;i++{
		wg.Add(1)

		go write()
	}

	for i:=0; i<10;i++{
		wg.Add(1)
		go read()
	}

	wg.Wait()

	end:= time.Now()

	fmt.Println(end.Sub(start))
}
