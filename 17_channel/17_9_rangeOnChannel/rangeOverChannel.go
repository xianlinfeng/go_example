package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		/* This range iterates over each element as it’s received from queue.
		Because we closed the channel above, the iteration terminates after receiving the 2 elements. */
		fmt.Println(elem)
	}
}
