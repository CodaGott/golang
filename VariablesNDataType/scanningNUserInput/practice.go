package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)


func main(){
	fmt.Println("Enter your name to get a greeting")
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var userInput =  scanner.Text()

	fmt.Println("Enter your age")
	var scanner2 = bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	var userInput2, _ = strconv.ParseInt(scanner2.Text(), 10, 64)

	fmt.Println("What number do you want to add to your age")
	var number = bufio.NewScanner(os.Stdin)
	number.Scan()
	var input, _ = strconv.ParseInt(number.Text(), 10 , 64)

	
	fmt.Println("Good morning", userInput , " and your age is" , userInput2 + input)
}