package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type StudentSlice []Student

func (s StudentSlice) Len() int {
	return len(s)
}

func (s StudentSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s StudentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	//切片排序
	s := []int{3, 5, 1, 6, 2, 7, 4}
	sort.Ints(s)
	fmt.Println(s)

	fmt.Println("----------------------------")
	//结构体切片排序
	var stuslice StudentSlice
	for i := 0; i < 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("学生:%d", rand.Intn(99)+1),
			Age:  rand.Intn(99) + 1,
		}
		stuslice = append(stuslice, stu)
	}
	fmt.Println("排序前------------")
	for _, student := range stuslice {
		fmt.Println(student)
	}

	fmt.Println("排序后------------")
	sort.Sort(stuslice)
	for _, student := range stuslice {
		fmt.Println(student)
	}
}
