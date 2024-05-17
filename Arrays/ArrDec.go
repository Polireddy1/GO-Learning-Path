package main

import "fmt"

func main() {
	// Declare and initialize an array using the shorthand syntax
	greetings := [2]string{"Bonjour", "Hola"}
	fmt.Println(greetings[0], greetings[1]) // Output individual elements
	fmt.Println(greetings)                  // Output the entire array

	// Using a composite literal to declare and initialize an array
	fibonacci := [...]int{0, 1, 1, 2, 3, 5, 8}
	fmt.Println(fibonacci)

	// Declare an array with predefined size and then initialize each element separately
	var cities [3]string
	cities[0] = "Tokyo"
	cities[1] = "Paris"
	cities[2] = "New York"
	fmt.Println(cities)
}