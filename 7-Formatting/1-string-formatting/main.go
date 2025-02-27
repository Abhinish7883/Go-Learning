/*
	- "fmt.Printf" is used for formatted output.


Specifier	Description					Example

	%s		String						"Hello" → fmt.Printf("%s", "Hello") → Hello

	%q		Quoted string				 "Go" → fmt.Printf("%q", "Go") → "Go"

	%x		Hexadecimal representation
			(lowercase)					"hello" → fmt.Printf("%x", "hello") → 68656c6c6f

	%X		Hexadecimal (uppercase)		"hello" → fmt.Printf("%X", "hello") → 68656C6C6F

	%v		Default format				fmt.Printf("%v", "hello") → hello

	%+v		Include field names (for structs)

	%#v		Go-syntax representation	 "hello" → fmt.Printf("%#v", "hello") → "hello"

	%T		Type of variable			 "hello" → fmt.Printf("%T", "hello") → string

*/

package main

import "fmt"

func main() {
	str := "Go lang"

	fmt.Printf("String: %s\n", str)
	fmt.Printf("Quoted String: %q\n", str)
	fmt.Printf("Hexadecimal: %x\n", str)
	fmt.Printf("Upper Hexadecimal: %X\n", str)
	fmt.Printf("Go syntex: %#v\n", str)
	fmt.Printf("Type of Variable: %T\n", str)
}
