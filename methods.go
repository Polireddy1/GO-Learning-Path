// package main

// import "fmt"

// type square struct {
//     side int
// }

// func (s *square) area() int {
//     return s.side * s.side
// }

// func (s square) perim() int {
//     return 4*s.side
// }

// func main() {
//     s := square{side: 5}

//     fmt.Println("area: ", s.area())
//     fmt.Println("perim:", s.perim())

//     rp := &s
//     fmt.Println("area: ", rp.area())
//     fmt.Println("perim:", rp.perim())
// }

// package main
// import "fmt"

// func main(){
// 	var side int
// 	fmt.Println("Enter a value: ")
// 	fmt.Scanln(&side)

// 	square :=area(side)
// 	sp :=perimeter(side)
// 	fmt.Println("Area is :",square)
// 	fmt.Println("Perimeter is :",sp)
// }
// func area(side int) int{
// 	return side*side
// }
// func perimeter(side int) int{
// 	return 4*side
// }