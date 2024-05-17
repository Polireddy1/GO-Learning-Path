package main

import "fmt"

func main() {
    // Declaring and initializing a 2D array
    var matrix = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    // Accessing elements in a 2D array
    fmt.Println("Element at row 0, column 1:", matrix[0][1]) // Output: 2

    // Iterating over a 2D array
    fmt.Println("Iterating over 2D array:")
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("Element at [%d][%d]: %d\n", i, j, matrix[i][j])
        }
    }
}
