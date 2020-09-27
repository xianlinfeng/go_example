package main

import "fmt"

func fact(n int) int { // get the factorial
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func fib(n int) int { // get the fibonacci sequence
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fact(7))

	for i := 1; i <= 40; i++ {

		fmt.Print("  ", fib(i))
	}
}
