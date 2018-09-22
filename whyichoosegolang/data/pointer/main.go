package main

import (
	"fmt"
	"reflect"
)

func main() {
	// b := *getPointer()
	fmt.Println("Value is", getPointer(), reflect.ValueOf(getPointer()))
}

func getPointer() (myPointer *int) {
	a := 234
	return &a
}
