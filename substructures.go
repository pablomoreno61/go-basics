package main

import "fmt"

type Square struct {
	Center Point
	length int
}

type Point struct {
	X int
	Y int
}

func (square *Square) Move(dx int, dy int) {
	square.Center.X += dx
	square.Center.Y += dy
}

func (square Square) Area() int {
	return square.length * square.length
}

func main() {
	square, err := NewSquare(10, 10, 20)
	if err != nil {
		fmt.Printf("Square cannot be moved")
	}

	square.Move(5, 6)

	fmt.Printf("X position: %d\n", square.Center.X)
	fmt.Printf("Y position: %d\n", square.Center.Y)
	fmt.Printf("Square area: %d\n", square.Area())
}

func NewSquare(x int, y int, length int) (*Square, error) {
	point := Point{x, y}

	square := Square{point, length}

	return &square, nil
}