// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var number int
// 	fmt.Println("Enter a number: ")
// 	fmt.Scanln(&number)

// 	sum := sumOfDigits(number)
// 	fmt.Println("Sum of digits:", sum)
// }

// func sumOfDigits(number int) int {
// 	sum := 0
// 	numberStr := strconv.Itoa(number)
// 	for _, digitChar := range numberStr {
// 		digit, _ := strconv.Atoi(string(digitChar))
// 		sum += digit
// 	}
// 	return sum
// }
