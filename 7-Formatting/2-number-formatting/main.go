/*
	Specifier		Description						Example
		%d			Integer (base 10)				123 → fmt.Printf("%d", 123) → 123
		%b			Binary representation			5 → fmt.Printf("%b", 5) → 101
		%o			Octal representation			9 → fmt.Printf("%o", 9) → 11
		%x			Hexadecimal (lowercase)			255 → fmt.Printf("%x", 255) → ff
		%X			Hexadecimal (uppercase)			255 → fmt.Printf("%X", 255) → FF
		%f			Floating-point number			3.1415 → fmt.Printf("%f", 3.1415) → 3.141500
		%.2f		Floating-point with precision	3.1415 → fmt.Printf("%.2f", 3.1415) → 3.14
		%e			Scientific notation (lowercase)	1234.567 → fmt.Printf("%e", 1234.567) → 1.234567e+03
		%E			Scientific notation (uppercase)	1234.567 → fmt.Printf("%E", 1234.567) → 1.234567E+03

*/

package main

import "fmt"

func main() {
	nums := 123
	pi := 3.14555
	expo := 1234.567

	fmt.Printf("%d\n", nums)
	fmt.Printf("%b\n", nums)
	fmt.Printf("%o\n", nums)
	fmt.Printf("%x\n", nums)
	fmt.Printf("%X\n", nums)

	fmt.Printf("%f\n", pi)
	fmt.Printf("%.2f\n", pi)

	fmt.Printf("%e\n", expo)
	fmt.Printf("%E\n", expo)
}
