/*
	Type Assertion (data.(type)) is used in a switch to check the type.
	interface{} is used for dynamic type checking.
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {

	var data interface{} = 56.89

	switch v := data.(type) {
	case int:
		fmt.Println("int", v)
	case bool:
		fmt.Println("Bool", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow type: ", reflect.TypeOf(v), v)
	}
}
