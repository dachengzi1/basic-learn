package main

import (
	"fmt"
	"reflect"
)

//func test() {
//	panic("this is panic")
//}
func main() {
	defer fmt.Println(111)
	defer fmt.Println(recover())
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err, reflect.TypeOf(err))
		}

	}()
	panic("this is panic")
	//panic("this is panic2")
}
