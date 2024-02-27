package main

import "fmt"

func WriteData(intchan chan int) {
	for i := 1; i <= 50; i++ {
		intchan <- i
		fmt.Println("写入=", i)
	}
	close(intchan)
}

func ReadData(intchan chan int, exitchan chan bool) {
	for {
		v, ok := <-intchan
		if !ok {
			break
		}
		fmt.Println("读取=", v)
	}
	exitchan <- true
	close(exitchan)
}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool)
	go WriteData(intChan)
	go ReadData(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
