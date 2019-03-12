package main

import (
	"fmt"
	"strconv"
)

func main() {
	str, num := unnamedReturns(1, 2)
	fmt.Println(str)
	fmt.Println(num)

	fmt.Println(namedReturns(3, 4))
}

func unnamedReturns(number1, number2 int) (string, int) {
	return strconv.Itoa(number1), number2
}

// All expected parameters are ints
func namedReturns(randomNumbers ...int) (number1 int, number2 int) {
	number1 = randomNumbers[0]
	number2 = randomNumbers[1]
	return
}
