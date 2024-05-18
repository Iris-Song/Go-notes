package main

import (
	"flag"
	"fmt"
)

// Command-line flags are a common way to specify options for command-line programs.
// For example, in wc -l the -l is a command-line flag.

func main() {
	// Basic flag declarations are available for string, integer, and boolean options. Here we declare a string flag word with a default value "foo" and a short description.
	// This flag.String function returns a string pointer (not a string value); we’ll see how to use this pointer below.
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// It’s also possible to declare an option that uses an existing var declared elsewhere in the program.
	// Note that we need to pass in a pointer to the flag declaration function.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// go build command-line-flags.go
// ./command-line-flags -word=opt -numb=7 -fork -svar=flag
// ./command-line-flags -word=opt
// ./command-line-flags -word=opt a1 a2 a3
// ./command-line-flags -word=opt a1 a2 a3 -numb=7
// ./command-line-flags -h
