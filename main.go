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

import (
	"fmt"
	// "strconv"
)

func convpen(a int) int{
	b:=a*60
	return b
}
func practice(tier string) int{
	
	switch tier {
	case "basic":
		return convpen(123)
	case "adv":
		return convpen(456)
	case "more":
		return convpen(789)
	default:
		return 0
}
}
func main() {
	result:=practice("basic")
	fmt.Println(result) 
}

