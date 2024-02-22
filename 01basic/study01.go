package main

import "fmt"

func main() {
	name := "张三"
	age := 20
	fmt.Println(name, "\n", age)

	a := []int{1, 2, 3, 4}
	fmt.Print("\n", a)

	if age >= 18 {
		fmt.Print("\n成年了")
	} else {
		fmt.Print("\n未成年")
	}

	count := 10
	for i := 0; i < count; i++ {
		fmt.Print("\n", i)
	}
}
