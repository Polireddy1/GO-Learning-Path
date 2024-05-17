package main

import "fmt"

func main() {

	// Demonstrating operations on a multidimensional slice
	mslice := [][]int{{1, 2}, {3, 4}, {5, 6}}

	fmt.Println("Original Multidimensional Slice:", mslice)

	// Example: Incrementing each element by 1
	for i := range mslice {
		for j := range mslice[i] {
			mslice[i][j]++
		}
	}

	// Printing the manipulated multidimensional slice
	fmt.Println("Manipulated Multidimensional Slice:")
	for _, row := range mslice {
		fmt.Println(row)
	}
}
