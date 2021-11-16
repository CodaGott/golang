package main

import "fmt"

func name(name string){
	fmt.Println(name)
}

func add (num1, num2 int){
	fmt.Println(num1, "+", num2, "=",num1 + num2)
}

func main() {

	world := "World!"

	name(world)

	add(2, 4)

	fmt.Println("Hello,", world)
}