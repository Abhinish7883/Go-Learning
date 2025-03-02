/*
	Higher Order Function:
		- Passing Functions as Arguments.
*/

package main

import "fmt"

func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {

	add := func(x int, y int) int {
		return x + y
	}

	result := applyOperation(45, 78, add)
	fmt.Println("Sum: ", result)

}
