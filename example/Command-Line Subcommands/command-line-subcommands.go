package main

import (
	"flag"
	"fmt"
	"os"
)

// Command-line subcommands are a common idiom for command-line tools, especially when the commands are related.
// For example, go build and go get are two different subcommands of the go tool.

func main() {
	// We declare a subcommand using the NewFlagSet function, and proceed to define new flags specific for this subcommand.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		// If no subcommand is provided, print an error and exit.
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

// go build command-line-subcommands.go
// ./command-line-subcommands foo -enable -name=joe a1 a2
// ./command-line-subcommands bar -enable a1
