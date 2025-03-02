/*
	Defer: Delaying Function Execution

		- The defer keyword delays execution of a function until the surrounding function returns.
		- defer is useful for cleanup, like closing files.
*/

package main

import "fmt"

func delayFunction() {
	fmt.Println("Start")
	defer fmt.Print("delay execution.......")
	fmt.Println("End")
}

func main() {
	delayFunction()
}
