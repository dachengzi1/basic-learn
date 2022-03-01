package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	cont, err := ioutil.ReadFile("./main.go")
	if err !=nil{
		fmt.Println("open file err:", err)
		return
	}
	fmt.Println("open file data:", string(cont))


	//file, err := os.Open("./main.go")
	//if err != nil {
	//	fmt.Println("open file err:", err)
	//	return
	//}
	//defer file.Close()
	//tmp := make([]byte, 128)
	//var list []byte
	//for{
	//	n, err := file.Read(tmp)
	//	if err == io.EOF {//判断是否读完
	//		fmt.Println("read end")
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("read file err:", err)
	//		return
	//	}
	//	list = append(list, tmp[:n]...)
	//}
	//
	//fmt.Println(string(list))


}
