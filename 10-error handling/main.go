package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b <= 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func customErrorCheck(id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("invalid user ID: %d", id)
	}
	return "User 123", nil
}

var ErrNotFound = errors.New("item not found")

func findItem(id int) error {
	if id != 1 {
		return fmt.Errorf("findItem: %w", ErrNotFound) // Wrapping error
	}
	return nil
}

func main() {

	number, err := divide(3, 1)

	// error Check
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", number)
	}

	// custom error check
	user, err := customErrorCheck(-1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User :", user)
	}

	// Using errors.Is() and errors.As() for Error Wrapping
	error1 := findItem(2)
	if errors.Is(error1, ErrNotFound) { // Checking error type
		fmt.Println("Custom Error: Item not found")
	} else if error1 != nil {
		fmt.Println("Error:", error1)
	}

}
