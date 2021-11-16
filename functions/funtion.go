package main

import "fmt"

func name(name string){
	fmt.Println(name)
}

func main() {

	world := "World!"

	name(world)

	fmt.Println("Hello,", world)
}