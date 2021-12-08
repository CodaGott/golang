package main

import "fmt"

func main(){

	currentPopulation := 200.0

	growth := 1.0

	estimatedGrowth := 0.0

	for year := 0; year < 6; year++ {
		estimatedGrowth = estimatedGrowth + currentPopulation * growth
		fmt.Println("Estimated growth for year", year, "equals", estimatedGrowth)
	}
}