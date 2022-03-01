package main

import "fmt"

type Cat struct {
	name  string
	color string
}

type BlackCat struct {
	Cat
}

func base() *Cat {
	return &Cat{}
}

func newBlackCat(name string) *BlackCat {
	cat := &BlackCat{}
	cat.name = name
	return cat
}
func main() {
	c := newBlackCat("黑猫")
	fmt.Printf("cat:%#v", c)
}
