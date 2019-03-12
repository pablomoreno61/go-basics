package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(resulterIf(-3))
	fmt.Println(resulterIf(0))
	fmt.Println(resulterIf(5))

	fmt.Println(resulterIf2(2))

	fmt.Println(resulterSwitch(3))
}

func resulterIf(number1 int) string {
	var result string

	// number2 does not exist out of condition scope
	if number2 := number1 * 2; number2 < 0 {
		result = "Less then 5"
	} else if number2 == 0 {
		result = "Is equal to 0"
	} else {
		result = "Greater than 5"
	}

	return result
}

func resulterIf2(number1 int) string {
	var result string
	var number2 int

	if number2 = number1 * 2; number2 < 0 {
		result = "Less then 5"
	} else if number2 == 0 {
		result = "Is equal to 0"
	} else {
		result = "Greater than 5"
	}

	return result + " " + strconv.Itoa(number2)
}

func resulterSwitch(number int) string {
	var result string

	switch {
	case number < 0:
		result = "Less than 0"
	case number == 0:
		result = "Equal to 0"
	default:
		result = "Greater than 0"
	}

	var resultInt int
	switch localScope := result; localScope {
	case "Less than 0":
		resultInt = -1
		result = localScope + ": " + strconv.Itoa(resultInt)
	case "Equal to 0":
		resultInt = 0
		result += ": " + strconv.Itoa(resultInt)
	case "Greater than 0":
		resultInt = 1
		result += ": " + strconv.Itoa(resultInt)
	}

	return result
}