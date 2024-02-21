package main

import "fmt"

// 作业1: 输入时间，打印年月日
func printTime(year int, month int, day int) {
	if month > 12 || month < 1 {
		fmt.Println("输入的月份不正确")
	} else {
		fmt.Printf("%v年%v月%v日", year, month, day)
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			fmt.Println("本共有31天")
		case 4, 6, 9, 11:
			fmt.Println("本月共有30天")
		case 2:
			if year%4 == 0 && year%100 != 0 {
				fmt.Println("输入的是闰年，2月有29天")
			} else {
				fmt.Println("输入的是平年，2月只有28天")
			}
		}
	}
}

func main() {
	//作业1
	var (
		year  int
		month int
		day   int
	)
	fmt.Println("请输入年份:")
	fmt.Scan(&year)

	fmt.Println("请输入月份:")
	fmt.Scan(&month)

	fmt.Println("请输入天:")
	fmt.Scan(&day)

	printTime(year, month, day)
}
