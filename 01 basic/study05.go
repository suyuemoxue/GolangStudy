package main

import "fmt"

func add01() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包
func cal(base int) (func(int) int, func(int) int) {
	add := func(num1 int) int {
		base += num1
		return base
	}
	sub := func(num2 int) int {
		base -= num2
		return base
	}
	return add, sub
}

// 递归
func jieCheng(num1 int) int {
	if num1 == 1 {
		return num1
	} else {
		return num1 * jieCheng(num1-1)
	}
}

func init() {
	fmt.Println("init")
}

func main() {
	var ff = add01()
	re := ff(10)
	fmt.Println("re = ", re)

	f1, f2 := cal(10)
	fmt.Println(f1(2), f2(5))

	fmt.Println(jieCheng(5))

	//defer
	fmt.Println("start")
	defer fmt.Println("step1")
	defer fmt.Println("step2")
	defer fmt.Println("step3")
	fmt.Println("end")
}
