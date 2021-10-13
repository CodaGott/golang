package main

import "fmt"

func main(){
	fmt.Println("Welcome to Slices in Golang")

	var array[5] int = [5] int{1, 2, 3, 4, 5}
	var sliced[]int= array[1:3]
	fmt.Println(cap(sliced))

	fmt.Println("New Slice Coming")
	var new_slice[]int = []int {9, 8, 7, 6, 5}
	var updated_slice =append(new_slice, 30)
	fmt.Println("updated slice",updated_slice)

}