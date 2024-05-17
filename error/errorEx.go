// package main

// import "fmt"

// func main() {
// 	userCount, err := getUserCount()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("Count of users is:%d", userCount)
// }

// func getUserCount() (user int, error error) {
// 	if uCount := -4; uCount <= 0 {
// 		return 0, fmt.Errorf("Caught error in userCount block the value passed is: %d", uCount)
// 	} else {
// 		return uCount, nil
// 	}
// }
