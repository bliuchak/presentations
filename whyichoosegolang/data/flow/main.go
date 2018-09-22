package main

import (
	"errors"
	"fmt"
)

func main() {
	switch true {
	case 1 == 1:
		fmt.Println("Uh")
	default:
		fmt.Println("oh")
	}

	var err error
	if _, err = getResult(); err != nil {
		fmt.Println("Uh oh", err)
	}

	if _, err = getResult()
	log.Println(err)

	fmt.Println(err)
}

func getResult() (bool, error) {
	return true, errors.New("dummy")
}
