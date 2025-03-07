package main

import "fmt"

func main() {

	// Declaration and Initilization of slices
	var slices = []int{5, 6, 7, 7}
	fmt.Println(slices)

	// using make() slices

	var s = make([]int, 2, 5)
	fmt.Println(s, len(s), cap(s))

	// Slicing an Array
	arr := [5]int{1, 2, 3, 4, 5}
	sli := arr[1:4] // Creates a slice from index 1 to 3 (length 3, capacity 4)
	fmt.Println(sli)

	//  Slicing an Existing Slice
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[1:4] // Creates a slice from index 1 to 3 (length 3, capacity 4)
	fmt.Println(s2)

	// accessing and modifing
	s2[2] = 4555
	fmt.Println(s1[2], s2)

	// use append
	s3 := []int{1, 2, 3}
	s3 = append(s3, 4, 5) // Appends 4 and 5 to the slice
	fmt.Println(s3)       // Output: [1 2 3 4 5]

	// Copy of slice
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)   // Copies elements from src to dst
	fmt.Println(dst) // Output: [1 2 3]

	// nil and empty
	var nilSlice []int
	fmt.Println(nilSlice == nil)

	emptySlice := []int{}
	fmt.Println(emptySlice == nil)

	// Passing Slices to Functions
	smain := []int{1, 2, 3}
	modifySlice(smain)
	fmt.Println(smain) // Output: [100 2 3]

}

func modifySlice(s []int) {
	s[0] = 100
}
