/*
	only the for loop (no while or do-while loops).

	🔸 Basic for Loop Syntax

	for initialization; condition; post-statement {
    	Loop body
	}
*/

package main

import "fmt"

func main() {

	/*
		🔸 Basic for Loop Syntax:

		for initialization; condition; post-statement {
			Loop body
		}

	*/
	fmt.Println("For loop statement: ")
	count := 10

	for i := 0; i < count; i++ {
		fmt.Println("Iteration:", i)
	}

	/*
		Infinite Loop (for Without Condition):

		all loop conditions, it creates an infinite loop (like while(true) in other languages).

		🔹 Use break to exit infinite loops.
	*/

	fmt.Println("Infinite Loop (for Without Condition)")
	conster := 0

	for {
		fmt.Println("Infinite Loop", conster)
		conster++
		if conster == 5 {
			break // Exits loop
		}
	}

	/*
		🔹Looping with break and continue
			break → Stops the loop immediately.
			continue → Skips the current iteration and continues with the next.
	*/

	fmt.Println("🔹Looping with break and continue")

	for i := 1; i < count; i++ {
		// fmt.Println(i)
		if i == 5 {
			fmt.Println("Skipping 5")
			continue // skip Printing 5
		}

		fmt.Println("run", i)

		if i == 8 {
			fmt.Println("Stopping at 8")
			break // stop loop
		}
	}

	/*
		- for Loop Without Initialization & Post Statement:

			when incrementing inside the loop body.

	*/

	// Similar to while(i <= 5)
	fmt.Println("for Loop Without Initialization & Post Statement: ")
	init := 1

	for init <= 5 {
		fmt.Println("Iteration:", init)
		init++ // Increment inside loop body...
	}

	/*
		- Looping Over Arrays and Slices (Using range):

			The range keyword is used to iterate over slices, arrays, maps, and strings.
	*/

	numbers := []int{10, 20, 30, 40, 50}

	for index, value := range numbers {
		fmt.Println("Index: ", index, "Value: ", value)
	}

	fmt.Println("If you don’t need the index, use _ (blank identifier):")

	for _, val := range numbers {
		fmt.Println("Value: ", val)
	}

}
