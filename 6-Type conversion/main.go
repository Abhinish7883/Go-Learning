/*
	What is Type Conversion?
	- Type conversion (also called type casting) is the process of converting a variable from one data type to another.
	- Go does not perform implicit type conversions (like Python or JavaScript).

	Syntex :
		convertedValue := Type(originalValue)

	Supported Type Conversions

	From 			Type						To Type	Explicit Conversion Required?

	int → float64	✅ float64(x)				✅ Yes
	float64 → int	✅ int(y)					✅ Yes (loss of precision)
	string → int	❌ Not directly supported	❌ No (use strconv.Atoi())
	int → string	❌ Not directly supported	❌ No (use strconv.Itoa())
	rune → string	✅ string(rune)				✅ Yes
	byte → string	✅ string(byte)				✅ Yes
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Integer to Float Conversion
	fmt.Println("Integer to Float Conversion")
	var a int = 42
	var b float64 = float64(a) // Convert int to float

	fmt.Println("Integer:", a) // Output: 42
	fmt.Println("Float:", b)   // Output: 42.0

	// Float to Integer Conversion : Note -> int(f) does not round the number; it just removes the decimal.
	fmt.Println("Float to Integer Conversion")
	var f float64 = 42.99
	var i int = int(f) // Convert float to int (truncates decimal)

	fmt.Println("Float:", f)   // Output: 42.99
	fmt.Println("Integer:", i) // Output: 42

	/*
		Integer to String Conversion:
		- not allow direct conversion from int to string.
		- use strconv.Itoa() from the "strconv" package.
	*/
	fmt.Println("Integer to String Conversion:")
	var num int = 123
	str := strconv.Itoa(num) // Convert int to string

	fmt.Println("String:", str) // Output: "123"

	/*
		String to Integer Conversion
		- not support direct string-to-int conversion.
		-  Use "strconv.Atoi()".
	*/
	fmt.Println("String to Integer Conversion")
	str1 := "456"
	num, err := strconv.Atoi(str1)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Integer", num)
	}

	str2 := "45t"

	num2, err2 := strconv.Atoi(str2)

	if err2 != nil {
		fmt.Println("Error", err2)
	} else {
		fmt.Println("Integer", num2)
	}

	/*
		Byte (uint8) and Rune (int32) Conversion:
		Byte to String
	*/

	fmt.Println("Byte to String")
	var b1 byte = 'A'
	var s string = string(b1)

	fmt.Println("Byte:", b)   // Output: 65 (ASCII code of 'A')
	fmt.Println("String:", s) // Output: "A"

	/*
		Rune to String
	*/

	fmt.Println("Rune to String")
	var r rune = '🚀'
	var s1 string = string(r)

	fmt.Println("Rune:", r)    // Output: 128640 (Unicode code point of 🚀)
	fmt.Println("String:", s1) // Output: "🚀"

	/*
		String to Byte Slice ([]byte):
		- convert a string to a byte slice ([]byte) and vice versa.
	*/
	fmt.Println("String to Byte Slice ([]byte)")
	str3 := "GoLang"
	byteSlice := []byte(str3) // Convert string to byte slice
	str5 := string(byteSlice)

	fmt.Println("Byte Slice:", byteSlice) // Output: [71 111 76 97 110 103]
	fmt.Println("String", str5)

}
