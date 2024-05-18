package main

import "fmt"

func main() {
    // Creating and Initializing Arrays
    fmt.Println("Creating and Initializing Arrays:")

    // Declaring an array with explicit size and type
    var arr1 [5]int
    fmt.Println("arr1:", arr1)

    // Declaring and initializing an array with values
    arr2 := [5]int{1, 2, 3, 4, 5}
    fmt.Println("arr2:", arr2)

    // Shorter syntax for initialization without specifying size
    arr3 := [...]string{"Apple", "Banana", "Cherry"}
    fmt.Println("arr3:", arr3)

    // Accessing Elements in an Array
    fmt.Println("\nAccessing Elements:")
    fmt.Println("arr2[2]:", arr2[2])
    fmt.Println("arr3[1]:", arr3[1])

    // Updating Elements in an Array
    fmt.Println("\nUpdating Elements:")
    arr1[0] = 10
    arr2[2] = 30
    fmt.Println("arr1:", arr1)
    fmt.Println("arr2:", arr2)

    // Iterating Over an Array
    fmt.Println("\nIterating Over Arrays:")
    fmt.Println("Using for loop:")
    for i := 0; i < len(arr2); i++ {
        fmt.Printf("arr2[%d]: %d\n", i, arr2[i])
    }

    fmt.Println("Using for range loop:")
    for index, value := range arr3 {
        fmt.Printf("Index: %d, Value: %s\n", index, value)
    }

    // Array Length
    fmt.Println("\nArray Length:")
    fmt.Println("Length of arr1:", len(arr1))
    fmt.Println("Length of arr2:", len(arr2))
    fmt.Println("Length of arr3:", len(arr3))

    // Multidimensional Arrays
    fmt.Println("\nMultidimensional Arrays:")
    var matrix [2][3]int
    matrix[0][0] = 1
    matrix[0][1] = 2
    matrix[0][2] = 3
    matrix[1][0] = 4
    matrix[1][1] = 5
    matrix[1][2] = 6
    fmt.Println("matrix:", matrix)

    fmt.Println("Iterating over matrix:")
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("matrix[%d][%d]: %d\n", i, j, matrix[i][j])
        }
    }
}
