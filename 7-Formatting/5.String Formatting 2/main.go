/*
	fmt.Sprintf is used to format strings without printing them.
*/

package main

import "fmt"

func main() {
	name := "Alise"
	age := 23

	formattedString := fmt.Sprintf("Name: %s, Age: %d", name, age)

	fmt.Println(formattedString)
}
