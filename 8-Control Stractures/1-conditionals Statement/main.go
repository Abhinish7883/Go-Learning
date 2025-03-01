/*
	-  "if-else" statements for conditional execution.

	if condition {
      Code block executed if condition is true
	} else if anotherCondition {
    	Code block executed if anotherCondition is true
	} else {
    	Code block executed if all conditions are false
	}

	🔹 Logical Operators:

		&& → AND
		|| → OR
		! → NOT

*/

package main

import "fmt"

func main() {

	age := 78

	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age <= 60 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}

	// Variable declaration inside an if condition, making the scope limited to that block.

	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is old.")
	}

	// fmt.Println(num) 'num is not accesible here.'

}
