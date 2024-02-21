package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行的参数有:", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("arg[%v}=%v\n", i, v)
	}
}
