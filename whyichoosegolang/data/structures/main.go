package main

import "fmt"

type Person struct {
	Name string
}

type point struct {
	X, Y int
}

func main() {
	ps := Person{Name: "George"}
	fmt.Println(ps, ps.Name)
	p := point{1, 2}
	fmt.Println(p.X, p.Y)
	p.Y = 6
	fmt.Println(p.X, p.Y)
}
