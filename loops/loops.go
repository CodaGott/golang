package main

import "fmt"


func main(){

	num := 1

	for num < 20{
		fmt.Println("Hello number is", num, "here")
		num += 1
	}

	fmt.Println()

	for num2 := 0; num2 <= 10; num2 += 2{
		fmt.Println("From second loop", num2, )
	}
	
	fmt.Println()
	
	for i := 0; i < 1000; i++ {
		
		if i != 0 && i % 3 == 0 && i % 7 ==0 && i % 9 == 0 {
			fmt.Println("Here", i)
		}		
	}


}