package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func main() {
	emma := Student{name: "Emma", age: 23}

	fmt.Println(emma.age)
}