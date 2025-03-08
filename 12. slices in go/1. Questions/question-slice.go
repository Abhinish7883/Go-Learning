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

	/*
		Reverse a Slice

		Create a slice []int{1, 2, 3, 4, 5}.
		Write a function to reverse the slice in place (without creating a new slice).
		Print the reversed slice.
	*/

	reversedSlice := reverseSlice(mergeSlice)
	fmt.Println(mergeSlice, reversedSlice)

	/*
		Find the Maximum Value in a Slice

			Create a slice []int{45, 23, 67, 12, 89}.
			Write a function to find the maximum value in the slice.
			Print the maximum value.
	*/

	maxSlice := []int{45, 23, 67, 12, 89}
	maxValue := findMaxValue(maxSlice)
	fmt.Println(maxValue)

	/*
		Check if Two Slices Are Equal

		Create two slices: []int{1, 2, 3} and []int{1, 2, 3}.
		Write a function to check if the two slices are equal (same length and elements).
		Test the function with different slices.
	*/

	slice1, slice2 := []int{1, 2, 3}, []int{1, 2, 3}
	fmt.Println(isEqualSlice(slice1, slice2))

	/*
		Remove Duplicates from a Slice

			Create a slice []int{1, 2, 2, 3, 4, 4, 5}.
			Write a function to remove duplicate elements.
			Print the slice without duplicates.
	*/
	duplicatesSlice := []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Println(removeDuplicatesElement(duplicatesSlice))

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

func reverseSlice(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func findMaxValue(slice []int) int {
	maxValue := slice[0]
	for _, val := range slice {
		if val > maxValue {
			maxValue = val
		}
	}
	return maxValue
}

func isEqualSlice(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for index := range slice1 {
		if slice1[index] != slice2[index] {
			return false
		}
	}
	return true
}

func removeDuplicatesElement(slice []int) []int {
	var newSlice []int
	for _, val := range slice {
		if includes(newSlice, val) {
			continue
		}
		newSlice = append(newSlice, val)
	}
	return newSlice
}

func includes(slice []int, element int) bool {
	for _, val := range slice {
		if val == element {
			return true
		}
	}
	return false
}

// Need to learn about slice module.
