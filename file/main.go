package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile1() {
	cont, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	fmt.Println("open file data:", string(cont))
}

func readFile2() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("read file err:", file)
		panic(err.Error())
	}
	defer file.Close()

	temp := make([]byte, 128)
	var list []byte
	for {
		n, err := file.Read(temp)
		if err == io.EOF {
			fmt.Println("read end")
			break
		}
		if err != nil {
			fmt.Println("read file err:", err)
			break
		}
		list = append(list, temp[0:n]...)
	}
	fmt.Println("file str:", string(list))
	fmt.Println("over")
}

//方法3：缓冲读取（如果文件比较大的情况下建议是缓冲读取）

func readFile3() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("read file err:", file)
		panic(err.Error())
	}
	defer file.Close()

	var list []byte
	reader := bufio.NewReader(file)
	for {
		b, _, err := reader.ReadLine()

		if err == io.EOF {
			fmt.Println("read end")
			break
		}
		if err != nil {
			fmt.Println("read file err:", err)
			break
		}
		list = append(list, b...)
	}

	fmt.Println("file str:", string(list))

	fmt.Println("over")

}
func main() {
	//readFile1()
	//readFile2()
	readFile3()
}
