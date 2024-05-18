package main

import (
	"fmt"
	// "time"
	// "maps"
)

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println(i * i)
	}
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	t := 4
	for t < 7 {
		fmt.Println(t * t)
		t++
	}
	numbers := [5]int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}
}
