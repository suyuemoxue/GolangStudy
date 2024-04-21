package main

import "fmt"

const (
	B = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
)

func test() int {
	i := 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2", i)
	}()
	return i
}

func main() {
	fmt.Println("return", test())
	fmt.Println(B, KiB, MiB, GiB, TiB, PiB, EiB)
}
