package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {

	size := 10
	arr := make([]string, size)
	for i := range arr {
		arr[i] = gofakeit.Word()
	}

	fmt.Println("Array:", arr)
	fmt.Println("Length:", len(arr))
	fmt.Printf("Type: %T\n", arr)
}