package main

import "fmt"

func main() {
	fmt.Println("All Even num from 0 to 20: ")
	for i := range 20 {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
