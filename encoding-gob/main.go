package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

//gob 格式： 自描述的二进制序列

func writeFile() {
	file, err := os.Create("./output.gob")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer file.Close()

	str := "hello word"
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(str)
	if err != nil {
		fmt.Println("编码错误", err.Error())
		return
	} else {
		fmt.Println("编码成功")
	}
}

func readFile() {
	file, err := os.Open("output.gob")
	if err != nil {
		fmt.Println("文件读取失败", err.Error())
		return
	}
	defer file.Close()

	dec := gob.NewDecoder(file)
	str := ""
	err = dec.Decode(&str)
	if err != nil {
		fmt.Println("文件读取失败", err.Error())
	} else {
		fmt.Println("文件读取成功: ", str)
	}
}
func main() {
	fmt.Println("开始创建文件。。。。")
	writeFile()
	time.Sleep(time.Second)
	fmt.Println("开始读文件。。。")
	readFile()
}
