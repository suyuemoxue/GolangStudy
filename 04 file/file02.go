package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//写文件
	filePath := "E:/file01.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	write := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		write.WriteString("你好\n")
	}
	write.Flush()
}
