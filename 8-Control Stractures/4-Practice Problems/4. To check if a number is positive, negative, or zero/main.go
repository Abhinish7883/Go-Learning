package main

import "fmt"

func main() {
	num := 78

	switch {
	case num > 0:
		fmt.Println("Positive Number.")
	case num < 0:
		fmt.Println("negitive Number.")
	default:
		fmt.Println("Zero..")
	}
}
