package main

import "fmt"

type persion struct {
	name string
	age  int
}

func printMsg(msg *struct{
	name string
	age int
})  {
	fmt.Printf("printMsg, type:%T, value: %v\n", msg, msg)
}

func main() {
	var p1 = persion{
		name : "张三",
		age : 21,
	}
	//p1.name = "张三"
	//p1.age = 21
	p2 := new(persion)
	p2.name = "张四"
	p2.age = 22
	p3 := &persion{}
	p3.name = "张五"
	p3.age = 23
	fmt.Printf("type:%T, value: %v\n", p1, p1)
	fmt.Printf("type:%T, value: %v\n", p2, p2)
	fmt.Printf("type:%T, value: %v\n", p3, p3)


	msg := &struct {
		name string
		age int
	}{
		name: "张6",
		age: 24,
	}
	printMsg(msg)
}