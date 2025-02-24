package main

import "fmt"

/*
	- Constants are immutable values that cannot be changed once assigned.
    - Declared using const (unlike var, constants do not use :=).
*/

const Pi = 3.14159
const Greeting = "Hello, Go!"

func main() {

	fmt.Println("PI:", Pi)
	fmt.Println(Greeting)
}
