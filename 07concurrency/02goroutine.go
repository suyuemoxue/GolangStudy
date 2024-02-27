package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	lock  sync.Mutex
)

func test01(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock() //加锁
	myMap[n] = res
	lock.Unlock() //解锁
}

func main() {
	for i := 1; i <= 20; i++ {
		go test01(i)
	}

	time.Sleep(time.Second * 10)

	for k, v := range myMap {
		fmt.Printf("map[%v] = %v\n", k, v)
	}
}
