package main

import (
	"fmt"
	"strconv"
	"sync"
)

//var obj = make(map[string]int)
var wg = sync.WaitGroup{}

//func set(key string, value int) {
//	obj[key] = value
//}
//
//func get(key string) int {
//	value := obj[key]
//	return value
//}

var obj = sync.Map{}

func main() {

	for i := 1; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			obj.Store(key, i)
			//value, _ := obj.Load(key)
			//fmt.Printf("key is: %s, value is:%d\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
	obj.Delete("4")

	obj.Range(func(key, value interface{}) bool {
		fmt.Printf("key is: %s, value is:%d\n", key, value)
		return true
	})

}
