package main

import (
    "fmt"
    "strings"
)

func main() {

    myMap := map[string]int{"apple": 5, "banana": 3, "cherry": 8}

    // 116. Map - for range over a map
    fmt.Println("Iterating over the map:")
    for key, value := range myMap {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }

    // 117. Map - delete element
    delete(myMap, "banana")
    fmt.Println("Map after deleting 'banana':", myMap)

    // 118. Map - comma ok idiom
    value, ok := myMap["banana"]
    if ok {
        fmt.Println("Value of 'banana':", value)
    } else {
        fmt.Println("'banana' not found in the map.")
    }

    // 119. Map - counting words in a book
    text := "hello world hello hello hello world"
    wordCount := make(map[string]int)
    words := strings.Fields(text)
    for _, word := range words {
        wordCount[word]++
    }
    fmt.Println("Word count map:", wordCount)
}
