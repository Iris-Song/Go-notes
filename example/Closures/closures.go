package main

// Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	println(nextInt())
	println(nextInt())
	println(nextInt())

	newInts := intSeq()
	println(newInts())
}

// 1
// 2
// 3
// 1
