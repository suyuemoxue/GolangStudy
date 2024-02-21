package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()
	//reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dstFile.Close()
	//writer := bufio.NewWriter(dstFile)
	return io.Copy(dstFile, srcFile)
}

func main() {
	filePath01 := "E:/file01.txt"
	filePath02 := "D:/file01.txt"
	CopyFile(filePath02, filePath01) //拷贝文件
}
