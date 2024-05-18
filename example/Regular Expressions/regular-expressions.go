package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))
	// The Submatch variants include information about both the whole-pattern matches and the submatches within those matches.
	// For example this will return information for both p([a-z]+)ch and ([a-z]+).
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
