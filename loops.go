package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(loop())

	foreachLoop()

	indexedForeachLoop()
}

func loop() string {
	var number2 int

	// number2 does not exist out of loop scope
	for number1 := 0; number1 < 10; number1++ {
		number2 = number1
		fmt.Println("number: ", number1)
	}

	return "final number: " + strconv.Itoa(number2)
}

func foreachLoop() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	for i := range colors {
		fmt.Println(colors[i])
	}
}

func indexedForeachLoop() {
	colors := make(map[string]string)

	colors["Red"] = "FF0000"
	colors["Green"] = "00FF00"
	colors["Bluen"] = "0000FF"

	for color, hex := range colors {
		fmt.Printf("Color %v has hexa code %v\n", color, hex)
	}
}