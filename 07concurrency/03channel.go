package main

import (
	"fmt"
)

type Monster struct {
	Name string
	Age  int
}

func main() {
	intChan := make(chan interface{}, 3)

	intChan <- 10
	intChan <- 20
	monster := Monster{
		Name: "皮卡丘",
		Age:  12,
	}
	intChan <- monster

	fmt.Println(len(intChan), cap(intChan))
	<-intChan
	<-intChan
	res := <-intChan
	fmt.Println(res)
	a := res.(Monster)
	fmt.Println(a.Name)
	close(intChan)
}
