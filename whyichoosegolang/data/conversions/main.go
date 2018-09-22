package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 2
	f := float64(i)
	u := uint(i)

	fmt.Println(i, reflect.TypeOf(i))
	fmt.Println(f, reflect.TypeOf(f))
	fmt.Println(u, reflect.TypeOf(u))
	fmt.Println(i + u)
}
