package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		data := make([]byte, 1024)
		//等待客户端发送信息
		//fmt.Println("等待客户端", conn.RemoteAddr().String(), "发送信息")
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("客户端已退出")
			return
		}
		//打印从客户端接收到的信息
		fmt.Print("客户端:", string(data[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("错误信息:", err)
		return
	}
	defer listen.Close()

	//循环等待客户端连接
	for {
		fmt.Println("等待连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("错误信息:", err)
		} else {
			fmt.Println(conn.RemoteAddr().String(), "已连接")
			go process(conn)
		}
	}
}
