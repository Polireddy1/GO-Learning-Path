package main

import (
	"fmt"
)

func main() {
	// Integer
	var age int = 30
	fmt.Println("Integer (var):", age)

	// Floating-point number
	var height float64 = 5.9
	fmt.Println("Float (var):", height)

	// String
	var name string = "Alice"
	fmt.Println("String (var):", name)

	// Boolean
	var isStudent bool = true
	fmt.Println("Boolean (var):", isStudent)

	// Constants
	const pi float64 = 3.14
	const greeting string = "Hello"
	const isGoAwesome bool = true
	fmt.Println("Constant (pi):", pi)
	fmt.Println("Constant (greeting):", greeting)
	fmt.Println("Constant (isGoAwesome):", isGoAwesome)

	// Type inference
	autoName := "Charlie"
	autoAge := 40
	autoHeight := 6.1
	autoIsStudent := false
	fmt.Println("Type Inference (autoName):", autoName)
	fmt.Println("Type Inference (autoAge):", autoAge)
	fmt.Println("Type Inference (autoHeight):", autoHeight)
	fmt.Println("Type Inference (autoIsStudent):", autoIsStudent)

	// Short variable declaration
	country := "Wonderland"
	population := 1000000
	averageIncome := 50000.75
	isDeveloped := true
	fmt.Println("Short Variable Declaration (country):", country)
	fmt.Println("Short Variable Declaration (population):", population)
	fmt.Println("Short Variable Declaration (averageIncome):", averageIncome)
	fmt.Println("Short Variable Declaration (isDeveloped):", isDeveloped)

	// Multiple variable declaration
	var (
		city     string = "Metropolis"
		zipCode  int    = 12345
		currency string = "Dollar"
	)
	fmt.Println("Multiple Variable Declaration (city):", city)
	fmt.Println("Multiple Variable Declaration (zipCode):", zipCode)
	fmt.Println("Multiple Variable Declaration (currency):", currency)

	// Zero values
	var (
		zeroInt    int
		zeroFloat  float64
		zeroString string
		zeroBool   bool
	)
	fmt.Println("Zero Values (zeroInt):", zeroInt)
	fmt.Println("Zero Values (zeroFloat):", zeroFloat)
	fmt.Println("Zero Values (zeroString):", zeroString)
	fmt.Println("Zero Values (zeroBool):", zeroBool)
}
