package main

import (
	"fmt"
	"unicode/utf8"
)

// A Go string is a read-only slice of bytes.
// The language and the standard library treat strings specially - as containers of text encoded in UTF-8.
// In other languages, strings are made of “characters”. In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.
// https://go.dev/blog/strings
func main() {

	// s is a string assigned a literal value representing the word “hello” in the Thai language.
	// Go string literals are UTF-8 encoded text
	const s = "สวัสดี"

	// Since strings are equivalent to []byte,
	// this will produce the length of the raw bytes stored within.
	fmt.Println("Len", len(s))
	// This loop generates the hex values of all the bytes that constitute the code points in s.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
