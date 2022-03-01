package main

import (
	"fmt"
	"math/rand"
	"time"
)

var letterBytes = []rune(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`)

func RandStringRune(num int) string {
	arr := make([]rune, num)
	for k := range arr {
		n := rand.Intn(len(letterBytes))
		fmt.Println("rand num:", n)
		arr[k] = letterBytes[n]
	}
	return string(arr)
}

var letterStr = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterStr[rand.Intn(len(letterStr))]
	}
	return string(b)
}

func main() {
	//初始化随机数的资源库, 如果不执行这行, 不管运行多少次都返回同样的值
	rand.Seed(time.Now().UnixNano())
	str1 := RandStringRune(4)
	fmt.Println("str1:", str1)

	str2 := RandStringBytes(4)
	fmt.Println("str2:", str2)
}
