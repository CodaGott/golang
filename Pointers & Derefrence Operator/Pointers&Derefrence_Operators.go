package main

import "fmt"

func changeValue(str *string){
	*str = "changed!"
	fmt.Println()
}

func changeValue2(str string){
	str = "changed!"
}

func main(){
	
	toChange := "Hello"
	fmt.Println(toChange, "Here")
	changeValue(&toChange)
	fmt.Println(toChange)
	
	x := 7
	y := &x

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x)
}