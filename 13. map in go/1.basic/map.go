package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----------Map Creation----------------")
	// map initilization using map
	usersAges := make(map[string]int)
	usersAges["Abhinish"] = 23 //  add new val
	usersAges["Badal"] = 23
	fmt.Println("usersAges: ", usersAges)

	// usig Map Literals
	users := map[string]string{
		"firstName": "Abhinish",
		"lastName":  "Tiwari",
	}
	users["firstName"] = "Badal" // update value
	fmt.Println("Users: ", users, len(users))

	// retrieving values
	age := usersAges["Abhinish"]
	fmt.Println(age)

	// check is a key Exists
	isKeyExit(usersAges, "Abhinish") // for found
	isKeyExit(usersAges, "")         // for not found

	// Deleting a key
	delete(usersAges, "Abhinish ")
	fmt.Println(usersAges)

	// iteratrion  over Map
	for key, val := range usersAges {
		fmt.Println(key, ":", val)
	}

	// Nested Maps
	studentGrades := map[string]map[string]int{
		"Abhinish": {"Math": 95, "science": 83},
		"Badal":    {"Math": 92, "science": 78},
	}
	fmt.Println("Student Greade", studentGrades)
	fmt.Println("Abhinish math marks: ", studentGrades["Abhinish"]["Math"])

	// Maps With structs
	type Person struct {
		Name string
		Age  int
	}

	people := map[string]Person{
		"user1": {"Abhinish", 23},
		"user2": {"Badal", 23},
	}

	fmt.Println(people)

	// Copy Map

	original := map[string]int{"Alice": 25}
	copyMap := original

	copyMap["Alice"] = 30 // Modifies the original map!

	fmt.Println(original["Alice"]) // Output: 30

	clone := make(map[string]int)
	for key, value := range original {
		clone[key] = value
	}

	fmt.Println("Clone:", clone)
}

func isKeyExit(user map[string]int, key string) {
	age, exist := user[key]
	if exist {
		fmt.Println("Age Fount: ", age)
	} else {
		fmt.Println("Key don't exist")
	}
}
