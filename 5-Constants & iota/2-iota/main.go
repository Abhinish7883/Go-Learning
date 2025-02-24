package main

import "fmt"

/*
	- iota is a special constant used to create sequential numbers.
	- Its value starts at 0 and increments automatically.
*/

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Use _ to skip a value.
const (
	_        = iota // Skips 0
	Silver          // 1
	Gold            // 2
	Platinum        // 3
)

func main() {

	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)

	fmt.Println("Use _ to skip a value.")

	fmt.Println("Silver:", Silver)
	fmt.Println("Gold:", Gold)
	fmt.Println("Platinum:", Platinum)
}
