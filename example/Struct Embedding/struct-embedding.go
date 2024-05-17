package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// A container embeds a base. An embedding looks like a field without a name.
type container struct {
	base
	str string
}

func main() {
	// When creating structs with literals, we have to initialize the embedding explicitly;
	// here the embedded type serves as the field name.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// We can access the base’s fields directly on co, e.g. co.num.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	// Alternatively, we can spell out the full path using the embedded type name
	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
