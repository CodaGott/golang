package main

import "fmt"

func main() {

	test := func()  {
		fmt.Println("Going")
	}

	test()
	
	hello := sayHello
	hello()

	func ()  {
		fmt.Println("Hello")
	}()
}

func sayHello(){
	fmt.Println("When called, I say Hello!")
}


