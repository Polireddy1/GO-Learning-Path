package main

import "fmt"

func main() {
    var numbers = [5]int{1, 2, 3, 4, 5}

    // Using a traditional for loop
    fmt.Println("Iterating with traditional for loop:")
    for i := 0; i < len(numbers); i++ {
        fmt.Println(numbers[i])
    }

    // Using the range keyword
    fmt.Println("Iterating with range:")
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
