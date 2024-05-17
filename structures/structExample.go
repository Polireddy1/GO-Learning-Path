// package main

// import "fmt"

// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// 	Gender    Gender
// }
// type Gender struct {
// 	Male   bool
// 	Female bool
// }

// type Truck struct {
// 	bedSize int
// 	car     Car
// }

// type Car struct {
// 	make  string
// 	model string
// }

// func main() {
// 	p1 := Person{}
// 	p1.FirstName = "Shreeraj"
// 	p1.LastName = "Shetty"
// 	p1.Age = 22
// 	p1.Gender.Female = true

// 	fmt.Println("First Name:", p1.FirstName)
// 	fmt.Println("Last Name:", p1.LastName)
// 	fmt.Println("Age:", p1.Age)
// 	fmt.Println("Gender:", p1.Gender.Female)

// 	lanesTruck := Truck{
// 		bedSize: 10,
// 		car: Car{
// 			make:  "toyota",
// 			model: "camry",
// 		},
// 	}

// 	fmt.Println(lanesTruck.bedSize, lanesTruck.car.make)

// }
