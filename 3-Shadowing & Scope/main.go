package main

import "fmt"

/*
	Shadowing:
		1. When a variable declared in an inner scope (local variable) has the same name as a variable in an outer scope (global or parent scope).
		2. The inner variable "shadows" (hides) the outer variable within that specific block.
*/

/*
	Shadowing a Global Variable
*/
var message = "I am a global variable."

func main() {

	fmt.Println(message) // Print: I am a global variable.

	message := "I am a local variable." // Shadowing global variable
	fmt.Println(message)                // Prints: I am a local variable

	{
		message := "I am inside a block" // Shadowing the local variable
		fmt.Println(message)             // Prints: I am inside a block
	}

	fmt.Println(message) // Prints: I am a local variable (Block shadowing is gone)

}
