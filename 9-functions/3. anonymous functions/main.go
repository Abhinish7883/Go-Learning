/*
	Anonymous Function:
		- Functions Without a Name
		- don’t have a name and are defined inside another function.

*/

package main

import "fmt"

func main() {
	add := func(a int, b int) int {
		return a + b
	}

	sum := add(5, 7)
	fmt.Println("Sum: ", sum)
}
