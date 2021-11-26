package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func main() {
	emma := Student{name: "Emma",grades: []int{90, 90, 99}, age: 31}

	fmt.Println(emma.age)
}