package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40, 50}

	fmt.Println("Original Slice:", slice)

	// Iterating over a slice using for range
	fmt.Println("Iterating over slice:")
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Accessing value by index
	if len(slice) > 0 {
		fmt.Printf("Value at index 2: %d\n", slice[2])
	}

	// Appending to a slice
	slice = append(slice, 60)
	fmt.Println("Slice after appending 60:", slice)

	// Slicing a slice
	subSlice := slice[1:4] // elements at index 1 to 3
	fmt.Println("Sub-slice from index 1 to 3:", subSlice)

	// Deleting an element from a slice - removing the element at index 1
	slice = append(slice[:1], slice[2:]...)
	fmt.Println("Slice after deleting element at index 1:", slice)
}
