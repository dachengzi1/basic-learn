package main

import (
	"fmt"
	"time"
)

func main() {
	//验证timer只能执行一次
	t := time.NewTimer(time.Second)
	defer t.Stop()

	go func() {
		for {
			<-t.C
			fmt.Println("时间到：timer只能执行一次")
		}

	}()
	time.Sleep(4 * time.Second)

	//停止定时器

	//t2 := time.NewTimer(time.Second)
	//go func() {
	//	<-t2.C
	//	fmt.Println("t2 时间到：timer只能执行一次")
	//}()
	////time.Sleep(time.Second * 2)
	//
	//b := t2.Stop()
	//if b {
	//	fmt.Println("定时器已经停止")
	//}

	//重置定时器

	//t3:=time.NewTimer(time.Second)
	//t3.Reset(4 * time.Second)
	//fmt.Println(time.Now())
	//fmt.Println(<-t3.C)

	//重复执行
	//
	//ticker := time.NewTicker(time.Second)
	//i := 0
	//go func() {
	//	for {
	//		i++
	//		fmt.Println(<-ticker.C, i)
	//		if i == 5 {
	//			ticker.Stop()
	//			fmt.Println("停止")
	//
	//		}
	//	}
	//}()
	//
	//for {
	//
	//}

}
