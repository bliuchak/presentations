package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Done")
	}()
	fmt.Println("Working...")
}
