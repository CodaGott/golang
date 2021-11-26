package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s *Student) getAvarage() float32{


	avarage := 0

	for i := 0; i < len(s.grades); i++{
		avarage = avarage + s.grades[i]
	}

	return float32(avarage)/ float32(len(s.grades));
}

func main() {
	emma := Student{name: "Emma",grades: []int{90, 90, 99}, age: 31}

	fmt.Println(emma.age)
	fmt.Println(emma.getAvarage())
}