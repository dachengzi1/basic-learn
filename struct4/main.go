package main

import "fmt"

type Address struct {
	mobile   string
	province string
}

type User struct {
	name    string
	age     int
	address Address
}

func main() {
	p := User{
		name: "张三",
		age:  20,
		address: Address{
			mobile:   "11111",
			province: "北京",
		},
	}
	fmt.Printf("用户信息:%#v", p)
}
