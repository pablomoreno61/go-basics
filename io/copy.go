package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read")

	hash := md5.New()

	// Read given "r" string and apply md5 hasher
	if _, err := io.Copy(hash, r); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", hash.Sum(nil))
}
