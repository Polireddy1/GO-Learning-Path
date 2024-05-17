package main

import (
	"errors"
	"fmt"
)

func getCoords() (int, int) {
	var x int = 4
	var y int = 0
	return x, y
}

func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Cant divie by zero")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}

func main() {
	if length := -4; length < 0 {
		fmt.Println("Email is invalid")
	} else {
		fmt.Println("Email is valid")
	}

	a, b := getCoords()
	fmt.Println(a, b)

	a, b, err := calculator(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("MUlt:%d division:%d", a, b)
	}

}
