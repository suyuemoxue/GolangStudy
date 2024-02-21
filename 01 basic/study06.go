package main

import (
	"fmt"
	"time"
)

// 指针
func main() {
	a := [3]int{1, 2, 3}
	var pa [3]*int
	for i := 0; i < len(a); i++ {
		pa[i] = &a[i]
	}
	fmt.Println(pa)
	for i, _ := range pa {
		fmt.Println(*pa[i])
	}

	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
}
