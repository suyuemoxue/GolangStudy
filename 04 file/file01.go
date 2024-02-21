package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//第一种方式，带缓冲，适合大文件
	file, err := os.Open("E:/testGo.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("file = ", file)

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		fmt.Print(str)
		if err == io.EOF { //io.EOF表示文件的结尾
			break
		}
	}

	fmt.Println("\n---------------------------------------")

	//第二种方式，不带缓冲、一次性读取，适合小文件
	file01, err01 := os.ReadFile("E:/testGo.txt")
	if err01 != nil {
		fmt.Println(err01)
	}
	fmt.Println(string(file01))
}
