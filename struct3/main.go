package main

import "fmt"

type Persion struct {
	Name string
	Age  int
}

func NewPersion(name string, age int) *Persion {
	return &Persion{
		Name: name,
		Age:  age,
	}
}

func (p Persion) Dream() {
	fmt.Printf("%d岁的%s，有一个梦想!\n", p.Age, p.Name)
}

func (p * Persion)SetAge( age int)  {
	p.Age = age
}

func (p Persion)SetAge2( age int)  {
	p.Age = age
}

//1.需要修改接收者中的值
//2.接收者是拷贝代价比较大的大对象
//3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
func main() {
	p := NewPersion("张三", 10)
	p.Dream()
	p.SetAge(20)
	p.Dream()
	p.SetAge2(30)
	p.Dream()
}
