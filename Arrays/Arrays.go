package main

import "fmt"

func main() {
	var arr [5] int
	arr[1] = 10
	arr[0] = 100
	arr[3] = 20

	array := []int {4,5,6}


	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("Sum",sum)

	fmt.Println(arr, array)
}