package main

import (
	"fmt"
	// "time"
	"maps"
)

func main() {
	map1 := map[string]int{"eng": 96, "mat": 92}

	for key, value := range map1 {
		fmt.Println(key, ":", value)
	}
	fmt.Println(map1)

	map2 := make(map[string]int)
	// for key,value:=range map2{
	//   fmt.Println(key,":",value)
	// }
	map2["soc"] = 78
	fmt.Println(map2)

	fmt.Println(len(map1))

	value, found := map1["eng"]
	fmt.Println(value, found)
	value1, found1 := map1["hin"]
	fmt.Println(value1, found1)
	map1["eng"] = 100
	fmt.Println(map1)

	delete(map1, "mat")
	fmt.Println(map1)

	// map1=make(map[string]int)
	// fmt.Println(map1)

	if maps.Equal(map1, map2) {
		fmt.Println("both are equal")
	}

}
