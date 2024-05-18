package main

import "fmt"

func main() {
	// Creating and initializing maps
	fmt.Println("Creating and Initializing Maps:")

	// Empty map with make
	var map1 map[string]int = make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	fmt.Println("Map1:", map1)

	// Map literal with initial values
	map2 := map[string]string{
		"first":  "Alice",
		"second": "Bob",
	}
	fmt.Println("Map2:", map2)

	// Map with explicit type
	var map3 map[int]string
	map3 = map[int]string{
		1: "One",
		2: "Two",
	}
	fmt.Println("Map3:", map3)

	// Inserting elements into a map
	fmt.Println("\nInserting Elements:")
	map1["three"] = 3
	map2["third"] = "Charlie"
	map3[3] = "Three"
	fmt.Println("Map1:", map1)
	fmt.Println("Map2:", map2)
	fmt.Println("Map3:", map3)

	// Updating elements in a map
	fmt.Println("\nUpdating Elements:")
	map1["one"] = 10
	map2["first"] = "Alice Updated"
	map3[1] = "One Updated"
	fmt.Println("Map1:", map1)
	fmt.Println("Map2:", map2)
	fmt.Println("Map3:", map3)

	// Accessing elements in a map
	fmt.Println("\nAccessing Elements:")
	fmt.Println("map1['two']:", map1["two"])
	fmt.Println("map2['second']:", map2["second"])
	fmt.Println("map3[2]:", map3[2])

	// Checking existence of keys
	fmt.Println("\nChecking Existence of Keys:")
	val, exists := map1["four"]
	if exists {
		fmt.Println("map1['four'] exists with value:", val)
	} else {
		fmt.Println("map1['four'] does not exist")
	}

	// Deleting elements from a map
	fmt.Println("\nDeleting Elements:")
	delete(map1, "three")
	delete(map2, "third")
	delete(map3, 3)
	fmt.Println("Map1:", map1)
	fmt.Println("Map2:", map2)
	fmt.Println("Map3:", map3)

	// Iterating over a map
	fmt.Println("\nIterating Over Maps:")
	fmt.Println("Map1:")
	for key, value := range map1 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	fmt.Println("Map2:")
	for key, value := range map2 {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
	fmt.Println("Map3:")
	for key, value := range map3 {
		fmt.Printf("Key: %d, Value: %s\n", key, value)
	}

	// Length of a map
	fmt.Println("\nLength of Maps:")
	fmt.Println("Length of Map1:", len(map1))
	fmt.Println("Length of Map2:", len(map2))
	fmt.Println("Length of Map3:", len(map3))
}
