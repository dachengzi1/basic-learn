package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	//bytes.Buffer：	//https://www.yisu.com/zixun/451631.html

	//1。是golang标准库中的缓存区，具有读写方法和可变大小的字节存储功能

	//2。直接定义一个 Buffer 变量，而不用初始化

	//3。bytes.Buffer自动扩容：当向缓冲区写入数据时，首先会检查当前容量是否满足需求，如果不满足分三种情况处理

	//  1）当前内置缓冲区切片buf为空，且写入数据量小于bootstrap的大小（64字节），则bootstrap作为buf
	//
	//  2）当前未读数据长度+新写入数据长度小于等于缓冲区容量的1/2，则挪动数据（将未读的数据放到已读数据位置）
	//
	//  3）以上条件不满足，只能重新分配切片，容量设定为2*cap(b.buf) + n，即两倍原来的缓冲区容量+写入数据量大小

	//4。string() 方法： 是将字节数组强制转换string 【申请空间 + 拷贝】

	//strings.Builder

	//1。write:将字节数组append到原数组上

	//  1）当原切片长度小于1024时，新切片的容量会直接翻倍。而当原切片的容量大于等于1024时，会反复地增加25%，直到新容量超过所需要的容量。

	//2。string() 方法：Builder只是指针的转换。

	var b3 bytes.Buffer
	b3.Write([]byte("Hello")) // 可以直接使用

	fmt.Println("b3:", b3.String(), b3.Len(), b3.Cap())

	b3.Write([]byte("go语言")) // 可以直接使用
	fmt.Println("b3:", b3.String(), b3.Len(), b3.Cap())

	var b4 strings.Builder
	b4.Write([]byte("Hello"))
	fmt.Println("b3:", b4.String(), b4.Len(), b4.Cap())

}
