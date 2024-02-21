package main

import "fmt"

// 作业3:输出100以内的素数，每行5个，并求和
func sushu() {
	var num1, num2, sum int
	for num1 = 2; num1 <= 100; num1++ {
		for num2 = 2; num2 < num1; num2++ {
			if num1%num2 == 0 {
				break
			}
		}
		if num2 > (num1 / num2) {
			sum += num1
			fmt.Print(num1)
		}
	}
	fmt.Println(sum)
}

// 作业4: 打印a~z，Z~A
func pri() {
	for i := 97; i < 123; i++ {
		str := string(i)
		fmt.Print(str)
	}
	fmt.Println()
	for i := 90; i > 64; i-- {
		str := string(i)
		fmt.Print(str)
	}
}

func main() {
	sushu()
	pri()
}
