package main

import (
	"fmt"
	"math/rand"
)

// 作业2:生成随机数，猜
func ran() int {
	ran := rand.Intn(100) + 1
	return ran
}

func guess(ran int) {
	var count int
	var num int

	for count = 1; count <= 10; count++ {
		fmt.Println("请猜数字")
		fmt.Scanln(&num)
		if num == ran {
			fmt.Printf("共用了%v次", count)
			if count == 1 {
				fmt.Println("天才")
			} else if count >= 2 && count <= 3 {
				fmt.Println("很聪明")
			} else if count >= 4 && count <= 9 {
				fmt.Println("一般般")
			} else if count == 10 {
				fmt.Println("可算猜对了")
			}
			break
		}
	}
	if count > 10 {
		fmt.Println("你输了")
	}
}

func main() {
	ran := ran()
	fmt.Println("随机数是:", ran)
	guess(ran)
}
