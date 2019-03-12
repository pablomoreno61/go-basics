package main

import (
	"fmt"
	"io"
	"os"
)

func createFile(fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	ln, err := io.WriteString(file, "Another written line")
	fmt.Printf("Number of characters written %v", ln)
}

func main() {
	createFile("resources/written.txt")
}
