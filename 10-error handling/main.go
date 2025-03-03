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

func main() {

	number, err := divide(3, 1)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", number)
	}

}
