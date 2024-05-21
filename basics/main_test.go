// package main

// import (

// 	"fmt"

// )

// // add function takes two integers and returns their sum

// func add(a int, b int) int {

// 	return a + b

// }

// func main() {

// 	// Example usage of the add function

// 	num1 := 3

// 	num2 := 5

// 	sum := add(num1, num2)

// 	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

// }

// package main

// import "testing"

// func Test_practice(t *testing.T) {
// 	type args struct {
// 		tier string
// 	}
// 	tests := []struct {
// 		name string
// 		// args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// package main

// import (

// 	"fmt"

// )

// // add function takes two integers and returns their sum

// func add(a int, b int) int {

// 	return a + b

// }

// func main() {

// 	// Example usage of the add function

// 	num1 := 3

// 	num2 := 5

// 	sum := add(num1, num2)

// 	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

// }

package main

import "testing"

func Test_practice(t *testing.T) {
	type args struct {
		tier string
	}
	tests := []struct {
		name string
		// args args
		want int
	}{
		// TODO: Add test cases.
		{"basic", 7380},
		{"adv", 27360},
		{"more", 47340},
		{"default", 0},
		{"basi", 0},
	}
	for _, tt := range tests {
		got := practice(tt.name)
		if got != tt.want {
			t.Errorf("practice() = %v, want %v", got, tt.want)
		}

	}
}
