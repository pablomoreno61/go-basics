package main

import (
	"fmt"
	"test/ship"
)

func main() {
	var titanic ship.Ship

	titanic.Name = "Titanic"
	// titanic.year = 1921
	ship.Stop()

	fmt.Println(titanic)
}