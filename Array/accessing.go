package main

import "fmt"

func main() {
    var numbers = [5]int{1, 2, 3, 4, 5}
    
    // Accessing elements
    fmt.Println("First element:", numbers[0])
    fmt.Println("Last element:", numbers[4])
    
    // Modifying elements
    numbers[0] = 10
    fmt.Println("Modified first element:", numbers[0])
}
