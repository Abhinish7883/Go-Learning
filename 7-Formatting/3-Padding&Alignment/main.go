/*

Specifier				Description					Example
	%10s			Right-align string (10 spaces)	"Go" → fmt.Printf("%10s", "Go") → " Go"
	%-10s			Left-align string (10 spaces)	"Go" → fmt.Printf("%-10s", "Go") → "Go "
	%010d			Zero-padding for integers		42 → fmt.Printf("%010d", 42) → 0000000042

*/

package main

import "fmt"

func main() {

	fmt.Printf("%-10s %-10s %-10s\n", "Name", "Age", "Country")
	fmt.Println("--------------------------------")
	fmt.Printf("%-10s %-10d %-10s\n", "Alice", 25, "USA")
	fmt.Printf("%-10s %-10d %-10s\n", "Bob", 30, "UK")
	fmt.Printf("%-10s %-10d %-10s\n", "Charlie", 28, "India")
}
