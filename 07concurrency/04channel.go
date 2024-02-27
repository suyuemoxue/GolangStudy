package main

import "fmt"

func main() {
	ch := make(chan int, 101)
	for i := 0; i < 100; i++ {
		ch <- i * 2
	}
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
