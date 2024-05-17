package main

import "fmt"

func main() {
	// Declaring empty map using make
	map1 := make(map[string]int)
	map1["apple"] = 10
	fmt.Println("Map 1 (using make):", map1)

	// Declare using map literals
	map2 := map[string]int{"banana": 20, "cherry": 30}
	fmt.Println("Map 2 (using literals):", map2)

	// empty map without make, will be nil and cannot be used until initialized
	var map3 map[string]int
	if map3 == nil {
		map3 = make(map[string]int)
		map3["date"] = 40
	}
	fmt.Println("Map 3 (initially nil, then initialized):", map3)

	// nested map
	map4 := make(map[string]map[string]int)
	map4["outerKey"] = make(map[string]int)
	map4["outerKey"]["innerKey"] = 50
	fmt.Println("Map 4 (nested map):", map4)
}
