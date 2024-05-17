package main

import "fmt"

func main() {
    //initialized to nil
    var mySlice1 []int

    // Using a slice literal to declare and initialize the slice
    mySlice2 := []int{1, 2, 3}

    // Using the make function to create a slice with specified length and default capacity equal to length
    mySlice3 := make([]int, 5) // Length and capacity are 5

    // Using the make function to create a slice with specified length and capacity
    mySliceWithCap := make([]int, 5, 10) // Length 5, capacity 10

    // Starting with a nil slice and appending items to it
    var mySlice []int
    mySlice = append(mySlice, 1, 2, 3)

    // Creating a slice as a subsection of an existing array
    baseArray := [5]int{1, 2, 3, 4, 5}
    mySlice4 := baseArray[1:4] // Creates a slice including elements 2, 3, and 4

    // Print statements to verify the contents of each slice (added for demonstration)
    fmt.Println("mySlice1:", mySlice1)
    fmt.Println("mySlice2:", mySlice2)
    fmt.Println("mySlice3:", mySlice3)
    fmt.Println("mySliceWithCap:", mySliceWithCap)
    fmt.Println("mySlice:", mySlice)
    fmt.Println("mySlice4:", mySlice4)
}
