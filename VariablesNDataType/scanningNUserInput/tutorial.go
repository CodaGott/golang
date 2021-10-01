package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// fmt.Println("Hello")
	// fmt.Println("Hello")
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the way you are born: ")
	scanner.Scan()
	var input, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years old at the end of 2020", 2020 - input)

}