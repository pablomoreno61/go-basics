package main

import (
    "fmt"
    "github.com/rubencougil/1-hello-world/hello"
)

func main() {
    var say string
    say = "Hello"

    world := "World"

    fmt.Println(say + " " + world)
    hello.SayHello()
}
