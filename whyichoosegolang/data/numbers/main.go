package main

import (
	"fmt"
)

func main() {
	num1 := 3         // int
	num2 := 3.2       // float64
	num3 := 3 + 4i    // complex128
	num4 := byte('a') // byte (alias for uint8)

	// Other types
	var u uint = 7       // uint (unsigned)
	var p float32 = 22.7 // 32-bit float

	fmt.Println(num1, num2, num3, num4)
	fmt.Println(u, p)
}
