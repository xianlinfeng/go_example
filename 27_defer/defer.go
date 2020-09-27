/* Defer is used to ensure that a function call is performed later in a program’s execution, usually for
purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages. */

package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("/tmp/defer.txt")
	defer closeFile(f) // 直到main函数完成之前，才会执行它
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	// It’s important to check for errors when closing a file, even in a deferred function.
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}