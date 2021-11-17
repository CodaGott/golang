package main

import "fmt"

func name(name string){
	fmt.Println(name)
}

func add (num1, num2 int){
	fmt.Println(num1, "+", num2, "=",num1 + num2)
}

func multiply(num1, num2 int) int{
	fmt.Println("Multiply values")
	return num1 * num2
}

func main() {


	test := func()  {
		fmt.Println("Going")
	}

	test()

	world := "World!"

	name(world)

	add(2, 4)
	answer := multiply(2, 4)
	fmt.Println("Multiply",answer)

	fmt.Println("Hello,", world)
}