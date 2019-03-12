package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("%s\n", t)

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("%v\n", tomorrow.Day())

	longFormat := "Monday, January 2, 2006"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))
}
