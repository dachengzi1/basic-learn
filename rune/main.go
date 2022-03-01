package main

import (
	"fmt"
	"unicode/utf8"
)

//rune: 是int32的别名，用来区分字符值和数字值

//golang 里string底层是byte数组，中文字符在unicode下占用2哥字节，在utf-8下占用3个字节，golang默认编码正好是utf-8

//字符编码的过程

//ASCII编码：
//1.最开始是美国信息交换的标准编码，用于不同计算机间的通信，后来被国际组织认定为国际标准，称为ISO 646 编码
//2.通常是一个字节存储一个字符
//unicode编码：
//1. 为每种语言的每个字符设定了统一且唯一的二进制编码，以满足跨语言，跨平台进行的文本转换和处理，
//2. unicode编码并未规定一个字符占用几个字节。 utf-8就出现了

//utf-8
//
func main() {
	str := "hello 编程"
	fmt.Println("str length", len(str))

	//以下两种都可以得到str的字符串长度
	fmt.Println("utf8.RuneCountInString:", utf8.RuneCountInString(str))
	fmt.Println("[]rune(str): ", len([]rune(str)))

}
