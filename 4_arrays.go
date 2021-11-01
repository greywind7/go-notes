package main

import "fmt"

func main() {
	// initialized to 0
	var x [5]int
	fmt.Println(x)

	// can be initialized
	var y [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(y)

	// shorthand
	z := [...]string{"Hello", "World", "bebop"}

	// traversing an array
	for i, v := range z {
		fmt.Println("Index", i, "Value", v)
	}
}
