package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.8:8888")
	if err != nil {
		fmt.Println("错误信息:", err)
		return
	}
	//fmt.Println(conn.RemoteAddr().String(), "已连接")
	for {
		//从键盘接收信息
		reader := bufio.NewReader(os.Stdin)
		//将接收到的信息转换成string类型
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("错误信息:", err)
		}
		//如果输入exit则退出
		if strings.Trim(str, "\r\n") == "exit" {
			break
		}
		//将信息转换成[]byte类型，并发送到服务器端
		conn.Write([]byte(str))
	}
	defer conn.Close()
}
