package main

import "fmt"

func main() {
	hello := sayHello
	hello()
}

func sayHello(){
	fmt.Println("When called, I say Hello!")
}