/*
	Implement a FizzBuzz program using a for loop:
		Print "Fizz" for multiples of 3.
		Print "Buzz" for multiples of 5.
		Print "FizzBuzz" for multiples of both 3 and 5.
		Otherwise, print the number.
*/

package main

import "fmt"

func main() {

	for i := 0; i <= 30; i++ {
		switch {
		case (i%3 == 0 && i%5 == 0):
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Printf("%d\n", i)
		}
	}

}
