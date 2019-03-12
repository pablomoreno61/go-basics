package main

import (
	"fmt"
	"time"
)

func say(s string, iterations int) {
	for i := 0; i < iterations; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world", 5)
	say("hello", 2)
}
