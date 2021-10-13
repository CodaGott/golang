package main

import "fmt"

func main(){
	a := []int{1, 2, 3, 4, 5, 2}

	fmt.Println("Looping through the arrays")
	for i := 0; i < len(a); i++{
		fmt.Println(a[i])
		
	}
		fmt.Println("Find Dubplicate")
	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element == element2{
				fmt.Println(element)
			}
		}
	}
}