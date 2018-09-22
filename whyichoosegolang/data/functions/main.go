package main

import (
	"errors"
	"fmt"
)

var x int

func main() {
	myfunc := func() (bool, error) {
		if x == 0 {
			return false, errors.New("x can't be equal to zero")
		}
		return x > 10000, nil
	}

	fmt.Println(myfunc())
	fmt.Println(hello())
}

func hello() string {
	return "hello world"
}
