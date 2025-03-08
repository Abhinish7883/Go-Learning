package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 70, 80)
	fmt.Println(slice)

	sliceNew := slice[1:3]
	fmt.Println(sliceNew, slice)

	slice[1] = 1000
	fmt.Println(slice)

	// copy a slice
	newSlice := slice[:]
	fmt.Println(newSlice)

	// Using a copy;
	newSlice2 := make([]int, len(newSlice))
	copy(newSlice2, newSlice)
	fmt.Println(newSlice2)

	/*
		Check if a Slice is Empty or Nil:

		- Create a nil slice and an empty slice.
		- Write a function to check if a slice is nil or empty.
		- Test the function with both slices.

	*/

	var nilSlice []int
	fmt.Println(nilSlice == nil)

	emptySlice := []int{}
	fmt.Println(emptySlice == nil)

	/*
		Remove an Element from a Slice

			Create a slice []int{1, 2, 3, 4, 5}.
			Write a function to remove the element at index 2.
			Print the updated slice.

	*/

	orignalSlice := []int{1, 2, 3, 4, 5}

	updatedSlice := removeElement(orignalSlice, 2)
	fmt.Println(orignalSlice, updatedSlice)

	/*
		Merge Two Slices

			Create two slices: []int{1, 2, 3} and []int{4, 5, 6}.
			Merge them into a single slice.
			Print the merged slice.

	*/

	mergeSlice := mergeTwoSlice(orignalSlice, updatedSlice)
	fmt.Println(mergeSlice)

	/*
		Filter a Slice

			Create a slice []int{10, 20, 30, 40, 50}.
			Write a function to filter out all elements greater than 25.
			Print the filtered slice.
	*/

	filterSlice := filterSlice(mergeSlice, 3)

	fmt.Println(filterSlice)

}

func removeElement(slice []int, index int) []int {
	if (index < 0) || index > len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func mergeTwoSlice(slice []int, slice2 []int) []int {
	return append(slice, slice2...)
}

func filterSlice(slices []int, limit int) []int {
	var filterSlices []int
	for _, value := range slices {
		if value > limit {
			filterSlices = append(filterSlices, value)
		}
	}

	return filterSlices
}
