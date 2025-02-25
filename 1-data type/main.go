/*
	- Go is a statically typed language, meaning every variable has a specific type that is determined at compile time.
	- Go provides a variety of built-in data types for handling different kinds of values.

	1. Boolean -> bool -> Ex: true/false
	2. Integer -> int, int8, ...., int64 -> Ex: 42, -100
	3. Unsigned Integer -> uint, uint8, uint16, uint32, uint64 -> Ex: 255, 10000
	4. Floating Point -> float32, float64 -> EX: 3.14, -0.001
	5. Complex Number -> complex64, complex12  -> Ex: 3 + 4i
	6. String -> string -> "Hello, Go!"
	7. Byte & Rune -> byte (alias for uint8), rune (alias for int32) -> EX: 'A', '💡'

*/

package main

import "fmt"

func main() {

	// Boolean
	var isGoFun bool = true
	fmt.Println("Is Go fun?", isGoFun) // Output: Is Go fun? true

	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false

	// Integer
	var a int = -10
	var b uint = 20

	fmt.Println("Integer:", a)  // Output: -10
	fmt.Println("Unsigned:", b) // Output: 20

	/*
		Floating-Point Numbers:
		- float32 : 32-bit ->  6-7 decimal places
		- float64: 64 -> 15-16 decimal places
	*/

	var pi float64 = 3.141592653589793
	fmt.Println("Pi:", pi) // Output: 3.141592653589793

	/*
		Complex Numbers (complex64, complex128)
		- A complex number has a real and imaginary part (a + bi).
		- Use the complex() function to create one.
	*/
	c := complex(3, 4)                // 3 + 4i
	fmt.Println("Complex number:", c) // Output: (3+4i)

	/*
		Strings (string)
		- sequence of UTF-8 characters.
		- Strings are immutable (cannot be changed once created).
	*/
	var name string = "GoLang"
	fmt.Println("Name:", name)        // Output: GoLang
	fmt.Println("Length:", len(name)) // Output: 6

	/*
		Byte (byte) and Rune (rune)
		- byte: Alias for uint8, represents a single ASCII character.
		- rune: Alias for int32, represents a Unicode character.
	*/
	var ch byte = 'A'
	var unicode rune = '🚀'

	fmt.Println("Byte (ASCII):", ch)        // Output: 65
	fmt.Println("Rune (Unicode):", unicode) // Output: 128640
}
