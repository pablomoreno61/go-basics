package main

import (
	"fmt"
	"reflect"
)

type Robot interface {
	PowerOn(model string, year string)
	PowerOff()
}

type Terminator struct {
	Year  string
	Model string
}

func (t *Terminator) PowerOn(model string, year string) {
	t.Model = model
	t.Year = year
}

func (t *Terminator) PowerOff() {
	t.Model = ""
	t.Year = ""
}

func structCasting(terminator *Terminator) {
	terminator.PowerOn("T800", "1984")

	fmt.Println("Struct Robot: ", terminator)
	fmt.Println("Struct Type: ", reflect.TypeOf(terminator))

	fmt.Println("Struct Model: ", terminator.Model)
	fmt.Println("Struct Year: ", terminator.Year)

	terminator.PowerOff()

	fmt.Println("Struct Model: ", terminator.Model)
	fmt.Println("Struct Year: ", terminator.Year)
}

func interfaceCasting(robot Robot) {
	robot.PowerOn("T800", "1984")

	fmt.Println("Interface Robot: ", robot)
	fmt.Println("Interface Type: ", reflect.TypeOf(robot))

	// Year property is not available here
	// fmt.Println(robot.Year)
}


func main() {
	terminator := &Terminator{}
	structCasting(terminator)
	interfaceCasting(terminator)

	numbers()
}

func numbers() {
	var number1 float64 = 42.13
	pointer1 := &number1
	fmt.Printf("Value of position %p is %v for type %T\n", pointer1, *pointer1, pointer1)

	var number2 *float64
	if number2 != nil {
		fmt.Println("number2 is: ", *number2)
	} else {
		number2 = &number1
		*number2 = *number2 * 10.00
		fmt.Println("value of number1: ", number1)
		fmt.Println("value of number2: ", *number2)

		number3 := float64(10)
		fmt.Printf("number3: %v", number3+*number2)
	}
}