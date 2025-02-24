package main

import "fmt"

func main() {
	// Declaring Variables: ->

	var name string = "Abhinish"
	age := 18
	var salary float64 = 50000.45
	var isEmployed bool = true

	fmt.Println("Name: ", name)
	fmt.Println("age: ", age)
	fmt.Println("salary: ", salary)
	fmt.Println("Is Employed:", isEmployed)

	/*
		Multiple variable decleratrions: ->
	*/
	var a, b, c int = 10, 20, 30
	fmt.Println("Multiple Variables: ", a, b, c)

	/*
		Multiple variable decleratrions of different types: ->
	*/
	var (
		ab       float64 = 45.890
		isActive         = true
	)
	fmt.Println("Different type variable decleration", ab, isActive)

	/*
		Type Conversion (Explicit Casting): ->
	*/
	var x int = 42
	var y float64 = float64(x)
	fmt.Println("Converted int io float:", y)

	/*
		String Manipulation: ->
	*/
	firstName := "Abhinish"
	lastName := "Tiwari"
	fullName := firstName + " " + lastName
	fmt.Println("Full Name:", fullName)

	fmt.Println("------------------Assignment:----------------------")

	/*
		1.Convert a "float64" value to an "int": ->
	*/
	pi := 3.1405467
	i := int(pi)
	fmt.Println("Converted float to int:", i)

	/*
		1. Formate String: ->
	*/

	fmt.Printf("Hi, My name is %s. I am %d years old.\n", fullName, age)

	/*
		TODO: Memory Allocation & Zero values: ->
	*/

	fmt.Println("Memory Allocation & Zero values: ")
	var str string
	var num int
	var decimal float64
	var isTrue bool

	fmt.Println("Default string:", str)      // ""
	fmt.Println("Default int:", num)         // 0
	fmt.Println("Default float64:", decimal) // 0.0
	fmt.Println("Default bool:", isTrue)     // false

}
