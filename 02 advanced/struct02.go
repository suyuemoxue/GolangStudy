package main

import "fmt"

type Dogs struct {
	name  string
	age   int
	color string
}

func printDogs(dogs *Dogs) {
	fmt.Println(dogs.name, "\t", dogs.age, "\t", dogs.color)
}

func main() {
	var dog01, dog02 Dogs

	dog01.name = "大黄"
	dog01.age = 5
	dog01.color = "yellow"

	dog02.name = "大黑"
	dog02.age = 7
	dog02.color = "black"

	printDogs(&dog01)
	printDogs(&dog02)

	var dogss *Dogs
	dogss = &dog01
	dogss.age = 3
	printDogs(dogss)
	printDogs(&dog01)
}
