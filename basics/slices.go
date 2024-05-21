// slice
package main

import (
	"fmt"
	// "time"
	// "maps"
)

func main() {

	a1 := [5]int{1, 2, 3, 4, 5}
	s1 := a1[:3]
	fmt.Println(s1)
	s1 = append(s1, 6, 7, 8, 9, 10)
	fmt.Println(s1)

	s2 := s1[:4]
	s3 := s1[5:]
	s4 := append(s2, s3...)
	fmt.Println(s4)

	s5 := make([]int, 6)
	num := copy(s5, s4)
	fmt.Println(s5)
	fmt.Println("number of elements copied:", num)

	start := 1
	end := 5
	s6 := make([]int, end-start)
	copy(s6, s4[start:end])
	fmt.Println("copied:   ", s6)

}
