package main

import "fmt"

func main() {

	// Declaration and Initilization
	var nums = [3]int{5, 4, 5}
	fmt.Println(nums)

	// shorthand Initilization
	nums2 := [3]int{2, 2, 2}
	fmt.Println(nums2)

	// using [...]
	nums3 := [...]int{4, 5}
	fmt.Println(nums3, nums3[0])

	// Iteration
	for i := 0; i < len(nums3); i++ {
		fmt.Printf("%d\n", nums3[i])
	}

	// Iterate using range

	for index, value := range nums2 {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Passing an Array Copy
	fmt.Println("Original:", nums2, "Modified: ", modify([3]int(nums2)))

	// Pass BY Refrence using Pointer
	// passByReference(&nums)
	fmt.Println(nums)

}

func modify(arr [3]int) (modifiedArr [3]int) {
	arr[0] = 45
	return arr
}

// func passByReference(arr *[3]int) {
// 	arr[2] = 456666
// }
