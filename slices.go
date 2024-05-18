package main

import "fmt"

func main() {
	// Creating and Initializing Slices
	fmt.Println("Creating and Initializing Slices:")

	// Declaring and initializing a slice
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice1:", slice1)

	// Creating a slice with make function
	slice2 := make([]string, 3)
	slice2[0] = "Apple"
	slice2[1] = "Banana"
	slice2[2] = "Cherry"
	fmt.Println("Slice2:", slice2)

	// Appending Elements to a Slice
	fmt.Println("\nAppending Elements to a Slice:")
	slice1 = append(slice1, 6, 7, 8)
	fmt.Println("Slice1 after append:", slice1)

	// Slicing a Slice
	fmt.Println("\nSlicing a Slice:")
	slicedSlice := slice1[2:5]
	fmt.Println("Sliced slice:", slicedSlice)

	// Copying Slices
	fmt.Println("\nCopying Slices:")
	slice3 := make([]int, len(slice1))
	copy(slice3, slice1)
	fmt.Println("Slice3 (copied from Slice1):", slice3)

	// Iterating Over a Slice
	fmt.Println("\nIterating Over a Slice:")
	for index, value := range slice2 {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// Length and Capacity of Slices
	fmt.Println("\nLength and Capacity of Slices:")
	fmt.Printf("Length of Slice1: %d, Capacity of Slice1: %d\n", len(slice1), cap(slice1))
	fmt.Printf("Length of Slice2: %d, Capacity of Slice2: %d\n", len(slice2), cap(slice2))
	fmt.Printf("Length of Slice3: %d, Capacity of Slice3: %d\n", len(slice3), cap(slice3))
}
