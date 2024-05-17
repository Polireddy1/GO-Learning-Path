// package main

// import (
// 	"errors"
// 	"fmt"
// )

// // type divideError struct {
// // 	dividendEle float64
// // }

// // // Error() implements builtin error interface.
// // func (d divideError) Error() string {
// // 	return fmt.Sprintf("Can not divide %v from 0", d.dividendEle)
// // }

// // func divide(dividend, divisor float64) (float64, error) {
// // 	if divisor == 0 {
// // 		return 0, divideError{dividendEle: dividend}
// // 	}
// // 	return dividend / divisor, nil
// // }
// // func main() {
// // 	res, err := divide(10, 0)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return
// // 	}
// // 	fmt.Printf("Result is %v", res)
// // }

// // Second method 2) using errors package

// func divide(dividend, divisor float64) (float64, error) {
// 	if divisor == 0 {
// 		return 0, errors.New("Can not divide from 0")
// 	}
// 	return dividend / divisor, nil
// }
// func main() {
// 	res, err := divide(10, 0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("Result is %v", res)
// }
