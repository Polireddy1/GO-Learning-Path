package main

import (
	"fmt"
	"time"
	// "maps"
)

func main() {
	num := [5]int{0, 1, 2, 3, 4}
	for i := 0; i < 5; i++ {
		if num[i] == 2 {
			fmt.Println(num[i])
			// break
			continue
		} else {
			fmt.Println(num[i])
		}
	}

	t := time.Now()
	fmt.Println(t)

	switch {
	case t.Hour() > 12:
		fmt.Println("evening")
	case time.Now().Hour() < 12:
		fmt.Println("Morningggg")
	}
}
