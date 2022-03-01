package main

import (
	"fmt"
	"runtime"
	"time"
)

func run1()  {
	for i:=0; i<100; i++{
		fmt.Println("run 1")
	}
}

func run2()  {
	for i:=0; i<100; i++{
		fmt.Println("run 2")
	}
}
func main()  {
	runtime.GOMAXPROCS(4)
	go run1()
	go run2()
	//runtime.Gosched()//停止当前gorountine ，执行其他的gorountine

	time.Sleep(time.Second)
}


