package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		_, err = conn.Write([]byte(input)) // 发送数据
		if err != nil {
			return
		}
		//buf := [512]byte{}
		//n, err := conn.Read(buf[:])
		//if err != nil {
		//	fmt.Println("recv failed, err:", err)
		//	return
		//}
		//fmt.Println(string(buf[:n]))
	}
}
