package main

import (
	"fmt"
	"os"
)

// Defer is used to ensure that a function call is performed later in a program’s execution,
// usually for purposes of cleanup.
// defer is often used where e.g. ensure and finally would be used in other languages.

// Suppose we wanted to create a file, write to it, and then close when we’re done.
// Here’s how we could do that with defer.
func main() {
	f := createFile("./tmp/defer.txt")
	defer closeFile(f)
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

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
