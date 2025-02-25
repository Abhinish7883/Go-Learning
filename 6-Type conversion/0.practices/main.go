package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Convert 100 (string) to an integer and multiply it by 2.")
	const str = "1000"
	integer, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Converted Value:", integer*2)
	}

	fmt.Println("")

	/*
		Convert a Unicode character into a rune and print its Unicode code point.
		Convert a string into a byte slice and print its ASCII values.
	*/

}
