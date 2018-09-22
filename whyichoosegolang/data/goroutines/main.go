package main

import "fmt"

func main() {
	// A "channel"
	ch := make(chan string)
	// Start concurrent routines
	go push("Moe", ch)
	go push("Larry", ch)
	go push("Curly", ch)
	// Read 3 results (Since our goroutines are concurrent,the order isn't guaranteed!)
	fmt.Println(<-ch, <-ch)
}

func push(name string, ch chan string) {
	msg := "Hey, " + name
	ch <- msg
}
