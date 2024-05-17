// package main

// import (
// 	"fmt"
// )

// func main() {
// 	myArr := []int{1, 2, 3, 4, 5}
// 	total := sum(myArr...)
// 	fmt.Printf("total is :%v", total)

// 	appendedArr := append(myArr, 0, 0, 0, 0)
// 	fmt.Println(appendedArr)

// 	appendedArr = append(myArr, myArr...)
// 	fmt.Println(appendedArr)
// }

// func sum(nums ...int) int {
// 	num := 0
// 	// for i := 0; i < len(nums); i++ {
// 	// 	num += nums[i]
// 	// }

// 	for i, num := range nums {
// 		num += nums[i]
// 	}
// 	return num
// }
