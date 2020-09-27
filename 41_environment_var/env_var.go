package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("FOO", "1") // set a env variable
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() { // Use os.Environ to list all key/value pairs in the environment.
		/* This returns a slice of strings in the form KEY=value.
		You can strings.SplitN them to get the key and value. Here we print all the keys. */
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}

/*
$ go run env_var.go
  FOO: 1
  BAR:
  ...

$ BAR=2 go run env_var.go
FOO: 1
BAR: 2
...
*/
