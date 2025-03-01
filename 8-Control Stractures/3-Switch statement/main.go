/*
	Switch Statement:

		- alternative to multiple "if-else" statements.
		- No need to "break" (it stops automatically).

		switch variable {
			case value1:
     			Code block
			case value2:
    			Code block
			default:
    			Default case (optional)
		}

*/

package main

import "fmt"

func main() {
	//  Switch statement....

	fmt.Println("Switch statement....")

	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}

	// Multiple Case matches in switch:
	fmt.Println("Multiple Case matches in switch:")

	fruit := "apple"

	switch fruit {
	case "apple", "banana":
		fmt.Println("It's a common fruit...")
	case "mango":
		fmt.Println("It's a tropical frruit")
	default:
		fmt.Println("Unknnown fruit...")

	}

	// switch true - Using Conditions in case

	age := 25

	switch {
	case age < 18:
		fmt.Println("Minor")
	case age >= 18 && age < 60:
		fmt.Println("Adult")
	default:
		fmt.Println("Senior Citizen..")

	}

}
