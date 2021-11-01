package main

import (
	"fmt"
)

func main() {
	// Simple variable declaration
	var x int = 8

	// can be inferred as well
	var y = x + 8

	// Short declaration can only be done in functions
	z := 69

	fmt.Println("X is", x, "Y is", y, "Z is "+string(rune(z)))

	// constants are basically static
	const a = 37
	const (
		hh = 88
		nz = 420
	)

	fmt.Printf("Constant a is %d\nHeil hitler %d\nNazi boi %d\n", a, hh, nz)

	// aliasing types
	type measure float32
	var length, breadth measure = 31.2, 67.43
	area := length * breadth

	fmt.Println("The length is ", length, "breadth is", breadth, "\nArea:", area)

	// Pointers exist
	ptr := &area
	fmt.Println("Area is stored at location", ptr)

	// new allocates a variable initialized with 0 and returns pointer to it
	randO := new(float32)
	*randO = 69.420
	fmt.Println("Random number has value", *randO, "and address", randO)
}
