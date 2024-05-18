package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Sorted:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Sorted:", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted:", s)
}
