package main

import (
	"fmt"
	"runtime"
	"time"
)

func f()  {
	for i:=0; i<10;i++{
		fmt.Println("f:",i)
	}
}

func f2()  {
	for i:=0; i<10;i++{
		fmt.Println("f2:",i)
	}
}
func main()  {
	runtime.GOMAXPROCS(2)
	go f()
	go f2()

	time.Sleep(time.Second)

}
