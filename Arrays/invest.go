package main

import "fmt"

func main () {

	arr := []int {3, 5, 1, 8, 9}

	priceBoughtAt := 7

	for i := 0; i < len(arr); i++{
		if arr[i] > priceBoughtAt{
			fmt.Println("Sold at", arr[i])
			break
		}else{
			fmt.Println("Nothing to sale")
		}
	}

}
