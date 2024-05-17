package main

import "fmt"

func main() {
    // Declaring an array of integers with a fixed size of 5
    var numbers [5]int
    fmt.Println("Empty array:", numbers)

    // Initializing an array at the time of declaration
    var initializedNumbers = [5]int{1, 2, 3, 4, 5}
    fmt.Println("Initialized array:", initializedNumbers)

    // Initializing an array without specifying the size
    inferredSizeArray := [...]int{10, 20, 30}
    fmt.Println("Array with inferred size:", inferredSizeArray)
}
