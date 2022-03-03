package main

import (
	"bytes"
	"fmt"
	"strings"
)

func plusConcat(n int, preStr string) string {
	str := ""
	for i := 0; i < n; i++ {
		str += preStr
	}
	return str
}

func printConcat(n int, preStr string) string {
	str := ""
	for i := 0; i < n; i++ {
		str = fmt.Sprintf("%s%s", str, preStr)
	}
	return str
}

//strings.Builder
//WriteString
//String
func stringConcat(n int, preStr string) string {
	var str strings.Builder
	for i := 0; i < n; i++ {
		str.WriteString(preStr)
	}
	return str.String()
}

func bufferConcat(n int, preStr string) string {
	var str bytes.Buffer
	for i := 0; i < n; i++ {
		str.WriteString(preStr)
	}
	return str.String()
}

func byteConcat(n int, preStr string) string {
	arr := make([]byte, 0)

	for i := 0; i < n; i++ {
		arr = append(arr, preStr...)
	}
	return string(arr)
}

func main() {
	str := plusConcat(10, "abv")
	printStr := printConcat(10, "abv")
	stringStr := stringConcat(10, "abv")
	bufferStr := bufferConcat(10, "abv")
	byteStr := byteConcat(10, "abv")
	fmt.Println(str)
	fmt.Println(printStr)
	fmt.Println(stringStr)
	fmt.Println("bufferConcat:", bufferStr)
	fmt.Println("byteStr:", byteStr)

	//var buf bytes.Buffer
	// buf.WriteString("adv")
	//buf.WriteString("adv2")
	//fmt.Println(buf.String())
	//date:=time.Now().Add(7 * 24 *time.Hour)
	//fmt.Println(date.Unix())

}
