package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	name string
	age  int
	city string
}

type student struct {
	name string
	age  int
}

func main() {

	type MyInt int //类型定义

	type youInt = int //类型别名

	var a MyInt
	var b youInt
	fmt.Printf("type of a: %T\n", a)
	fmt.Printf("type of b: %T", b)

	var p4 person
	fmt.Printf("p4=%#v\n", p4)     //p4=main.person{name:"", city:"", age:0}
	fmt.Println(unsafe.Sizeof(p4)) //40

	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		fmt.Println(&stu)
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}




}
