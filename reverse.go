package main

import (
	"fmt"
)

func reverseSlice(input []int) []int {
	length := len(input)
	reversed := make([]int, length)
	for i, v := range input {
		reversed[length-1-i] = v
	}
	return reversed
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	reversed := reverseSlice(slice)
	fmt.Println(reversed) // Output: [5, 4, 3, 2, 1]
}
