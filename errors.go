package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func readFile(fileName string) {
	_, err := ioutil.ReadFile(fileName)

	if err != nil {
		// open notexist.txt: no such file or directory
		fmt.Println(err.Error())
	}

	customError := errors.New("My first error")
	fmt.Println(customError)
}

func main() {
	// open resources/notexist.txt: no such file or directory
	readFile("resources/notexist.txt")
}

