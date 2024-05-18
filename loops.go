package main

import "fmt"

func main() {
	// Basic for loop (similar to traditional for loop)
	fmt.Println("Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Simulating a while loop using for
	fmt.Println("\nSimulated while loop:")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Simulating a do-while loop using for
	fmt.Println("\nSimulated do-while loop:")
	k := 0
	for {
		fmt.Println(k)
		k++
		if k >= 5 {
			break
		}
	}

	// For loop with range (slice)
	fmt.Println("\nFor loop with range (slice):")
	fruits := []string{"Apple", "Banana", "Cherry"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}

	// For loop with range (map)
	fmt.Println("\nFor loop with range (map):")
	grades := map[string]string{"Math": "A", "Science": "B+"}
	for subject, grade := range grades {
		fmt.Printf("Subject: %s, Grade: %s\n", subject, grade)
	}

	// Infinite for loop with break statement
	fmt.Println("\nInfinite for loop with break statement:")
	l := 0
	for {
		if l >= 5 {
			break
		}
		fmt.Println(l)
		l++
	}

	// For loop with continue statement
	fmt.Println("\nFor loop with continue statement:")
	for m := 0; m < 5; m++ {
		if m%2 == 0 {
			continue // Skip the even numbers
		}
		fmt.Println(m)
	}

	// If statements
	fmt.Println("\nIf statements:")
	age := 25
	if age < 18 {
		fmt.Println("Minor")
	} else if age >= 18 && age < 60 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior")
	}

	// If with short statement
	fmt.Println("\nIf with short statement:")
	if n := 10; n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}

	// Switch statements
	fmt.Println("\nSwitch statements:")
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Switch with multiple expressions
	fmt.Println("\nSwitch with multiple expressions:")
	grade := "B"
	switch grade {
	case "A", "A+":
		fmt.Println("Excellent")
	case "B", "B+":
		fmt.Println("Good")
	case "C":
		fmt.Println("Average")
	case "D":
		fmt.Println("Below Average")
	case "F":
		fmt.Println("Fail")
	default:
		fmt.Println("Invalid grade")
	}

	// Switch with no expression
	fmt.Println("\nSwitch with no expression:")
	number := 10
	switch {
	case number%2 == 0:
		fmt.Println(number, "is even")
	case number%2 != 0:
		fmt.Println(number, "is odd")
	}

	// Switch with fallthrough
	fmt.Println("\nSwitch with fallthrough:")
	switch n := 2; n {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other number")
	}
}
