package main

import "fmt"

func main() {

	messages := make(chan string) // make a channel

	go func() { messages <- "ping" }() // put a msg into the channel

	msg := <-messages // get a msg from the channel
	fmt.Println(msg)
}
