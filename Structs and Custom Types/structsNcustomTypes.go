package main

import "fmt"

type Point struct {
	x int32
	y int32
}

func main() {
	p1 := Point{1, 2}
	// var p2 Point = Point{-5, 6}

	fmt.Println(p1.x)
}