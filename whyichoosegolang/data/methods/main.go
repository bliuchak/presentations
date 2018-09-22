package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func main() {
	p := Point{1, 2}
	fmt.Println(p.Sum())
}

func (p Point) Sum() int {
	return p.X + p.Y
}
