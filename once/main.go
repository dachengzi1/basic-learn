package main

import (
	"fmt"
	"sync"
	"time"
)

type signle struct {
	i int32
}

var sig *signle
var once sync.Once

func getInstance() *signle {
	once.Do(func() {
		sig = &signle{i:1}
	})
	return sig
}
func main() {

	getInstance()
	getInstance()
	fmt.Println(*sig)
	time.Sleep(time.Second)

}
