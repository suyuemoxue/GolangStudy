package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x  int
	wg sync.WaitGroup
	rl sync.RWMutex
)

func read() {
	rl.RLock()
	time.Sleep(time.Millisecond)
	rl.RUnlock()
	wg.Done()
}

func write() {
	rl.RLock()
	x = x + 1
	time.Sleep(time.Millisecond * 10)
	rl.RUnlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
