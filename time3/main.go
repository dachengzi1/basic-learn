package main

import (
	"fmt"
	"time"
)

type req struct {
	Int int64
}

func main() {

	//创建一个周期性的定时器
	//ticker := time.NewTicker(3 * time.Second)
	//fmt.Println("当前时间为:", time.Now())
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	go func() {
		for {

			for {
				fmt.Println("start...")
				timer := time.NewTicker(3 * time.Second)
				select {
				case do := <-done:
					fmt.Println("do:", do)
					break
				case t := <-timer.C:
					fmt.Println("当前时间为:", t)
					break
				}

			}
		}
	}()

	for {
		time.Sleep(time.Second * 1)
	}
}
