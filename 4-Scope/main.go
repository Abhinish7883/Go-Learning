package main

import "fmt"

/*
	Scope in Go:

	1. Global Scope – Accessible throughout the package.
	2. Function Scope – Accessible only inside a function.
	3. Block Scope – Limited to {} block.
*/

var globalVar = "Accessible everywhere"

func main() {
	localVar := "Accessible only in main()"

	{
		bloclVar := "Accessible only inside this block"
		fmt.Println(bloclVar)
	}
	// fmt.Println(bloclVar) Error

	fmt.Println(localVar)
	fmt.Println(globalVar)

}
