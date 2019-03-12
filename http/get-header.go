package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	// Defer executes this command when the current function is finished
	defer resp.Body.Close()

	header := resp.Header.Get("Content-Type")

	return header, nil
}

func main() {

	cType, err := contentType("https://www.golang.org")

	if err != nil {
		fmt.Printf("Url cannot be read")
	}

	fmt.Printf("%s content-type found", cType)

}
