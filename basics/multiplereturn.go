package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {

	num1, num2 := vals()
	fmt.Println(num1)
	fmt.Println(num2)

	_, c := vals()
	fmt.Println(c)
}
