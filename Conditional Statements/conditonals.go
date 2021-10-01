package main

import "fmt"

func main (){

	name()

	name := "mike"

	if name == "dozie"{
		fmt.Println("Name is", name)
	}else if name == "mike"{
		fmt.Println("Name is mike")
	}else{
		fmt.Println("Name is not mike or dozie")
	}
}

func name(){
	age := 18

	if age == 18 || age > 18{
		fmt.Println("You can ride")
	}
}