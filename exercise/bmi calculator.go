package main

import (
	// "bufio"
	"fmt"
	// "os"
)

func main() {

	// scanner := bufio.NewScanner(os.Stdin)


	fmt.Println("Enter your weight in Ibs")
	// scanner.Scan()

	var weightInKG float64
	fmt.Scan(&weightInKG)


	heightInMeters := 61.32
	
	var bmi float64 =  weightInKG * 703.0 / (heightInMeters * heightInMeters)

	if bmi < 18.5 {
		fmt.Println("Underweight", bmi)
	}else if bmi >= 18.5 && bmi <= 24.9 {
		fmt.Println("Normal weight", bmi)
	}else if bmi >= 25 && bmi <= 29.9 {
		fmt.Println("Overweight", bmi)
	}else if bmi >= 30 {
		fmt.Println("Obese", bmi )
	}
}