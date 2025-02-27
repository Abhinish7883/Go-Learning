/*
	1. Create a table that lists student names, grades, and percentages.
	2. Format a currency value (e.g., Price: $1234.50).
	3. Convert an integer to binary, octal, and hexadecimal formats.
*/

package main

import "fmt"

func main() {

	fmt.Println("1. Create a table that lists student names, grades, and percentages.")

	fmt.Printf("%-10s %-10s %-10s\n", "Name", "Grades", "percentages")
	fmt.Println("--------------------------------")
	fmt.Printf("%-10s %-10d %-10d\n", "Alice", 25, 20)
	fmt.Printf("%-10s %-10d %-10d\n", "Bob", 30, 45)
	fmt.Printf("%-10s %-10d %-10d\n", "Charlie", 28, 70)

	fmt.Println("2. Format a currency value (e.g., Price: $1234.50).")

	strFormatting := fmt.Sprintf("Price:%s%f", "$", 1234.50)
	fmt.Println(strFormatting)

	fmt.Println("3. Convert an integer to binary, octal, and hexadecimal formats.")

	num := 456

	fmt.Printf("Binary: %b", num)
	fmt.Printf("Octal: %o", num)
	fmt.Printf("Hexadecimal: %x", num)

}
