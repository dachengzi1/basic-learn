package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(pre, cur string) (r int64, err error) {
	src, err := os.Open(pre)
	if err != nil {
		fmt.Println("open pre file:", err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(cur, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open cur file:", err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)

}
func main() {
	_, err:= copyFile("./main.go", "./test.go")
	if err!= nil{
		fmt.Println("copy file err:", err)
	}
	fmt.Println("copy file done")
}
