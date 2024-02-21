package main

import "fmt"

type Cats struct {
	name  string
	age   int
	color string
}

func printCat(cat Cats) {
	fmt.Println(cat.name, "\t", cat.age, "\t", cat.color)
}

func main() {
	var cat01, cat02 Cats

	cat01.name = "小白"
	cat01.age = 3
	cat01.color = "white"

	cat02.name = "小花"
	cat02.age = 10
	cat02.color = "black"

	fmt.Println("名字", "\t", "年龄", "\t", "颜色")
	printCat(cat01)
	printCat(cat02)
}
