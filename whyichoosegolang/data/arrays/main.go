package main

import (
	"fmt"
)

func main() {
	// array
	numbers := [...]int{0, 0, 0, 0, 0}
	strings := [2]string{"hello", "world"}

	// slice
	slice := []int{2, 3, 4}

	fmt.Println(numbers, len(numbers))
	fmt.Println(strings, len(strings))
	fmt.Println(slice, len(slice))
}
