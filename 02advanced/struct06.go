package main

import "fmt"

type Address struct {
	Province string
	City     string
}

type Person struct {
	Name string
	Age  int
	Address
}

func main() {
	p1 := Person{
		Name: "张三",
		Age:  22,
		Address: Address{
			Province: "安徽",
			City:     "合肥",
		},
	}
	fmt.Println(p1)
}
