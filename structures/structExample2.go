// // package main

// // import "fmt"

// // type rect struct {
// // 	width  int
// // 	height int
// // }

// // //functions for structures
// // func (r rect) area() int {
// // 	return r.width * r.height
// // }

// // func main() {
// // 	r := rect{
// // 		width:  6,
// // 		height: 10,
// // 	}

// // 	fmt.Println(r.area())

// // }

// package main

// import "fmt"

// // don't touch below this line

// type expense interface {
// 	cost() float64
// }

// type email struct {
// 	isSubscribed bool
// 	body         string
// 	toAddress    string
// }

// type sms struct {
// 	isSubscribed  bool
// 	body          string
// 	toPhoneNumber string
// }

// type invalid struct{}

// func (e email) cost() float64 {
// 	if !e.isSubscribed {
// 		return float64(len(e.body)) * .05
// 	}
// 	return float64(len(e.body)) * .01
// }

// func (s sms) cost() float64 {
// 	if !s.isSubscribed {
// 		return float64(len(s.body)) * .1
// 	}
// 	return float64(len(s.body)) * .03
// }

// func (i invalid) cost() float64 {
// 	return 0.0
// }

// func getExpenseReport(e expense) (string, float64) {
// 	// Type assertion to check if 'e' is of type 'email'   (type-assert e to an email struct)
// 	// em, ok := e.(email)
// 	// if ok {
// 	// 	// If 'e' is an 'email', return its 'toAddress' and calculated cost
// 	// 	return em.toAddress, em.cost()
// 	// }
// 	// sm, ok := e.(sms)
// 	// if ok {
// 	// 	return sm.toPhoneNumber, sm.cost()
// 	// }
// 	// return "", 0

// 	// //same above logic but with switch statement

// 	switch v := e.(type) {
// 	case email:
// 		return v.toAddress, v.cost()
// 	case sms:
// 		return v.toPhoneNumber, v.cost()
// 	default:
// 		return "", 0
// 	}
// }

// func main() {
// 	// Create an instance of 'email' structure
// 	e := email{
// 		isSubscribed: false,
// 		body:         "Hello World!",
// 		toAddress:    "example@example.com",
// 	}
// 	address, cost := getExpenseReport(e)
// 	fmt.Println("To Address:", address)
// 	fmt.Println("Cost:", cost)

// 	s := sms{
// 		isSubscribed:  true,
// 		body:          "Hello World!",
// 		toPhoneNumber: "8660095291",
// 	}
// 	phoneNumber, cost := getExpenseReport(s)
// 	fmt.Println("To PhoneNumber:", phoneNumber)
// 	fmt.Println("Cost:", cost)
// }
