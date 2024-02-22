package main

import "fmt"

func f1(score int) {
	if score >= 90 && score <= 100 {
		fmt.Println("优")
	} else if score >= 70 && score < 90 {
		fmt.Println("良")
	} else if score >= 60 && score < 70 {
		fmt.Println("及格")
	} else if score >= 0 && score < 60 {
		fmt.Println("不及格")
	} else {
		fmt.Println("成绩不合理")
	}
}

func f2() {
	var arr = [...]int{1, 2, 3}
	for key, value := range arr {
		fmt.Printf("key: %d\t value: %d \n", key, value)
		fmt.Printf("arr[%d] = %d \n", key, value)
	}
}

func f3() {
	ma := make(map[string]string)
	ma["name"] = "张三"
	ma["age"] = "20"
	for key, value := range ma {
		fmt.Printf("key: %s\tvalue: %s \n", key, value)
	}
}

func f4() {
	var s1 []int
	fmt.Println("s1 : ", s1)

	s2 := make([]int, 2)
	fmt.Println("s2 : ", s2)
	s2[0] = 1
	s2 = append(s2, 3)
	fmt.Println("s2 : ", s2)
	fmt.Println("s2的长度 = ", len(s2))
	fmt.Println("s2的容量 = ", cap(s2))

}

func f5() {
	arr1 := [...]int{1, 2, 3, 4, 5}

	slice1 := arr1[2:4]
	fmt.Println(slice1)

	slice2 := arr1[:3]
	fmt.Println(slice2)

	slice3 := arr1[3:]
	fmt.Println(slice3)

	for key, value := range slice2 {
		fmt.Printf("slice2[%d] = %d \n", key, value)
	}

}

func main() {
	var score int
	fmt.Println("请输入一个数字")
	fmt.Scan(&score)
	fmt.Println("成绩:", score)
	f1(score)

	f2()
	f3()
	f4()
	f5()
}
