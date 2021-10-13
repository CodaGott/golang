package main

import "fmt" 

func main(){
	var map_name map[string]int = map[string]int{
		"apple": 5, "orange": 4, "pear": 9,
	}
	fmt.Println("Getting all the values in the map")
	fmt.Println(map_name)
	fmt.Println("Getting just a value")
	fmt.Println(map_name["pear"])
	fmt.Println("Changing the value of an element and getting all the values in the map with a changed value for the pear")
	map_name["pear"] = 50
	fmt.Println(map_name)
}