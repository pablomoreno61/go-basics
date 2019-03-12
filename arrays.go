package main

import (
	"fmt"
	"sort"
)

func arrays() {
	// array: fixed size, , cannot be extended
	var cheeses[3] string

	cheeses[0] = "Grullere"
	cheeses[1] = "Feta"
	cheeses[2] = "Havarti" //==>  invalid array index 2

	fmt.Println(cheeses)
}

func slices() {
	// slice: fixed size, can be extended
	var cars = make([]string, 2)
	cars[0] = "Toyota"
	cars[1] = "Volvo"
	//cars[2] = "Volvo" //==> Index out of range

	// first argument to append must be slice
	cars = append(cars, "Seat", "Fiat")

	fmt.Println("Array: ", cars)
	sort.Strings(cars)
	fmt.Println("Array: ", cars)

	var motos = []string{"BMW", "Suzuki", "Ducati"}
	motos = append(motos[:len(motos) - 1], "Honda")
	fmt.Println("Unordered slice: ", motos)
	sort.Strings(motos)
	fmt.Println("Ordered slice: ", motos)
}

func maps() {
	// make allocates and initializes memory, without it, we will cause panic error, assignment to entry in nil map
	players := make(map[string]int)

	players["Mike"] = 8
	players["John"] = 10
	// cannot use "10" (type string) as type int in assignment
	//players["John"] = "10"
	fmt.Println("Map: ", players)

	delete(players, "Mike")

	players["Kurtz"] = 5
	players["Venkatesh"] = 5

	fmt.Println("Map after deleting Mike: ", players)
}

func main() {
	arrays()

	slices()

	maps()
}