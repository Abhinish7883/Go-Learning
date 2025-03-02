/*
 variadic functions that accept a variable number of arguments using "...".


*/

package main

import "fmt"

func calculateSum(nums ...int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

func main() {

	fmt.Println("Pass multiple argument")
	sum := calculateSum(6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 66, 6, 6, 0)
	fmt.Println(sum)

	fmt.Println("Pass a slice in variadic function: ")
	nums := []int{5, 6, 78, 7, 0, 6}
	sum2 := calculateSum(nums...)
	fmt.Println(sum2)

}
