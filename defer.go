package main

import "fmt"

func main() {
	var x int
	x = 1000

	// Deferring will store the value of x at the moment of its definition
	defer fmt.Println("x value when deferring: ", x)
	x++
	fmt.Println("x value: ", x)
}
