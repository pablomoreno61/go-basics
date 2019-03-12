package main

import (
    "fmt"
)

type Robot interface {
    PowerOn() string
    PowerOff() string
}

type Robocop struct {
    Name string
}

func (r Robocop) PowerOn() string {
    r.Name = "Alex"
    return "Protect and serve"
}

func (r Robocop) PowerOff() int {
    r.Name = "Alex"
    return 12
}

func activate (robot Robot) string {
    return robot.PowerOn()
}

func main() {
    robot := Robocop{}

    fmt.Println(robot)
    fmt.Println(robot.PowerOn())
    fmt.Println(robot.PowerOff())
}
