package main

import "fmt"

func main() {
	for col := 0; col < 5; col++ {
		for row := 0; row <= col; row++ {
			fmt.Printf("%c", '*')
		}
		fmt.Print("\n")
	}
}
