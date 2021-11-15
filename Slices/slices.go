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


	var new_slice2[]int = [] int {3, 34, 5, 4}
	fmt.Println("New slice 2 before append",new_slice2)
	
	append_to_slice := append(new_slice2, 10)

	fmt.Println("Updated new slice 2 with appended value", append_to_slice)

	fmt.Println("New slice 2 after append", new_slice2)

	make_slice := make([]int, 5)
	fmt.Println("The slice made with make method", make_slice)

}