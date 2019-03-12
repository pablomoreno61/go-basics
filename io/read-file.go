package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(fileName string) {
	file, err := ioutil.ReadFile(fileName)

	if err == nil {
		fmt.Println(string(file))
	} else {
		fmt.Println("Cannot read file: ", fileName)
	}
}

func main() {
	readFile("resources/test.txt")
}
