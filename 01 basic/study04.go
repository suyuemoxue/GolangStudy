package main

import "fmt"

func test1(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func test2() (name string, age int) {
	name = "tom"
	age = 21
	return name, age
}

func sayHello(name string) {
	fmt.Printf("hello,%s", name)
}

func test3(name string, f func(string)) {
	f(name)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func sub(num1 int, num2 int) int {
	return num1 - num2
}

func cal01(op string) func(int, int) int {
	switch op {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	fmt.Println(test1(3, 5))
	fmt.Println(test2())

	test3("jack", sayHello)

	ff := cal01("+")
	re := ff(2, 3)
	fmt.Println("re =", re)
}
