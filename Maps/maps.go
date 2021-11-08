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
	map_name["work"] = 10
	fmt.Println("Adding new Item to Map")
	fmt.Println()
	fmt.Println(map_name)
	fmt.Println()
	delete(map_name, "work")
	fmt.Println("Deleting the last added item")
	fmt.Println(map_name)

	vie, now := map_name["pear"]
	fmt.Println(vie, now)
	fmt.Println(map_name)
}