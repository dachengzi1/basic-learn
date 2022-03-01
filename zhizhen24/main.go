package main

import "fmt"

func main() {
	a := 1
	b := &a
	c := *b
	fmt.Printf("a:%d, %p, \n", a, &a)
	fmt.Printf("b:%d, %p, %T\n", *b, &b, *b)
	fmt.Printf("c:%d, %p\n", c, &c)

	//内存分配
	//new & make：
	//1 都是用来分配内存的
	//2 make适用于slice ，map ， channel 内存分配,返回是类型的本身
	//3 new 返回是类型的指针
	m := new(int)
	fmt.Printf("m: %T\n", m)

	obj := make(map[string]int, 10)
	obj["age"] = 1
	fmt.Print("obj info:", obj)
}
