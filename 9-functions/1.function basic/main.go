/*
	Funtion:
		- organize code into reusable blocks.
		- assigned to variables, passed as arguments, and returned from other functions.
		- defined using the "func" keyword.

*/

package main

import "fmt"

func main() {
	fmt.Println("First run main function....")

	fmt.Println("Funtion with no parameter and no return value.")
	sayHello()

	fmt.Println("funtion with parameters")
	sum(4, 6)

	fmt.Println("funtion with return value")
	sumVal := sumWithReturn(78, 89)
	fmt.Println("Sum Return: ", sumVal)

	fmt.Println("function with multiple return value")
	q, r := divide(45, 4)
	fmt.Printf("quotient: %d\n Remainder: %d\n", q, r)

	fmt.Println("Named Return Values: ")
	a, p := rectangleProps(5, 10)
	fmt.Println("Area: ", a, "Perimeter: ", p)

}

/*
	Funtion with no parameter and no return value.
*/

func sayHello() {
	fmt.Println("Hello Bro...")
}

/*
	funtion with parameters
*/

func sum(a int, b int) {
	fmt.Println("Sum:", a+b)
}

/*
	funtion with return value
*/

func sumWithReturn(a int, b int) int {
	return a + b
}

/*
	function with multiple return value
*/

func divide(devident int, divider int) (int, int) {
	quotient := devident / divider
	remainder := devident % divider
	return quotient, remainder
}

/*
	Named Return Values :
*/

func rectangleProps(len int, width int) (area int, perimeter int) {
	area = len * width
	perimeter = 2 * (len + width)
	return
}
