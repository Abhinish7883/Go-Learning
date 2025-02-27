/*

 */

package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func main() {

	p := Person{"Abhinish", 24}

	fmt.Printf("Default format: %v\n", p)
	fmt.Printf("Field names: %+v\n", p)
	fmt.Printf("Go syntax: %#v\n", p)

}
