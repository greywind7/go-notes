package main

import "fmt"

func main() {
	var x, y int
	// if else
	fmt.Print("Enter X and Y ")
	fmt.Scan(&x, &y)
	if x > y {
		println("X big")
	} else {
		fmt.Println("Y big")
	}

	// simple switch
	switch x {
	case 0:
		fmt.Println("Woops X is 0")
	default:
		fmt.Println("Yay X is not 0 :D")
	}

	// switch without an argument
	switch {
	case y < 0:
		fmt.Println("Y is smol :(")
	case y == 0:
		fmt.Println("Y is 0 :|")
	case y > 0:
		fmt.Println("Y is big :)")
	default:
		fmt.Println("Y is not a number")
	}

	// for loop rules
	for i := 0; i < 3; i++ {
		println("I am ", i)
	}

	j := 4
	for j > 0 {
		println("J is ", j)
		j--
	}

	for {
		var t int
		fmt.Print("Enter 0 to exit (or dont): ")
		fmt.Scan(&t)
		if t == 0 {
			println("Good boy :)")
			break
		} else {
			println("I said 0 ffs")
		}
	}
}
